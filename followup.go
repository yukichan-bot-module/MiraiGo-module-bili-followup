package followup

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/yukichan-bot-module/MiraiGo-module-bili-followup/internal/pkg"
	"gopkg.in/yaml.v2"
)

type updateNotification struct {
	GroupID int64
	Video   pkg.VideoRecord
}

var instance *followup
var logger = utils.GetModuleLogger("com.aimerneige.bili.followup")
var followConfig map[int64][]int64
var sleepMinutes int
var latestTimeMap map[int64]int64 = make(map[int64]int64)

type followup struct {
}

func init() {
	instance = &followup{}
	bot.RegisterModule(instance)
}

func (f *followup) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "com.aimerneige.bili.followup",
		Instance: instance,
	}
}

// Init 初始化过程
// 在此处可以进行 Module 的初始化配置
// 如配置读取
func (f *followup) Init() {
	path := config.GlobalConfig.GetString("aimerneige.followup.path")
	if path == "" {
		path = "./followup.yaml"
	}
	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &followConfig)
	if err != nil {
		logger.WithError(err).Error("Unable to read config file in %s", path)
	}
	sleepMinutes = config.GlobalConfig.GetInt("aimerneige.followup.sleep")
	for mid := range followConfig {
		vList, err := pkg.GetLatestVideoList(mid)
		if err != nil {
			logger.WithError(err).Errorf("Unable to get latest video list for mid: %s", mid)
		}
		if len(vList) <= 0 {
			logger.Error("No video found for mid: %s", mid)
			return
		}
		latestTime := vList[0].Created
		for _, item := range vList {
			if item.Created > latestTime {
				latestTime = item.Created
			}
		}
		latestTimeMap[mid] = latestTime
	}
}

// PostInit 第二次初始化
// 再次过程中可以进行跨 Module 的动作
// 如通用数据库等等
func (f *followup) PostInit() {
}

// Serve 注册服务函数部分
func (f *followup) Serve(b *bot.Bot) {
}

// Start 此函数会新开携程进行调用
// ```go
//
//	go exampleModule.Start()
//
// ```
// 可以利用此部分进行后台操作
// 如 http 服务器等等
func (f *followup) Start(b *bot.Bot) {
	for {
		for mid, groupList := range followConfig {
			vList, err := pkg.GetLatestVideoList(mid)
			if err != nil {
				logger.WithError(err).Error("Unable to get latest video list")
				continue
			}
			for _, video := range vList {
				if video.Created > latestTimeMap[mid] {
					// notify
					for _, groupID := range groupList {
						msg, err := getVideoMessage(b.QQClient, groupID, video)
						if err != nil {
							logger.WithError(err).Error("Unable to get video message for group: %s", groupID)
							continue
						}
						b.QQClient.SendGroupMessage(groupID, msg)
					}
					logger.Println("Notify:", video.Created, video.Title, video.Pic, video.URL)
				}
			}
			// update latestTime
			for _, item := range vList {
				if item.Created > latestTimeMap[mid] {
					latestTimeMap[mid] = item.Created
				}
			}
		}
		// sleep a while
		time.Sleep(time.Minute * time.Duration(sleepMinutes))
	}
}

// Stop 结束部分
// 一般调用此函数时，程序接收到 os.Interrupt 信号
// 即将退出
// 在此处应该释放相应的资源或者对状态进行保存
func (f *followup) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
}

func getVideoMessage(c *client.QQClient, groupID int64, video pkg.VideoRecord) (*message.SendingMessage, error) {
	picBytes, err := pkg.HTTPGetRequest(video.Pic, [][]string{})
	if err != nil {
		return nil, err
	}
	imgMsgElement, err := c.UploadGroupImage(groupID, bytes.NewReader(picBytes))
	if err != nil {
		return nil, err
	}
	msgString := fmt.Sprintf("%s更新视频了！\n标题：%s\n视频链接：%s", video.Author, video.Title, video.URL)
	textMsg := message.NewText(msgString)
	return message.NewSendingMessage().Append(imgMsgElement).Append(textMsg), nil
}
