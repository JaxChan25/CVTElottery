package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//联通测试
		/**
		* showdoc
		* @catalog 测试模块/连通性测试
		* @title ping
		* @description 查看后端是否连接正常
		* @method post
		* @url /api/v1/ping
		* @return {"code":0,"msg":"Pong"}
		* @remark 无备注
		 */
		v1.POST("ping", api.Ping)

		// 用户登录
		/**
		* showdoc
		* @catalog 用户模块
		* @title 抽奖用户注册
		* @description 抽奖用户的注册api
		* @method post
		* @url /api/v1/user/register
		* @param user_name 可选 string 用户名
		* @param password 必选 string 密码
		* @param password_confirm 必选 string 确认密码
		* @param real_name 必选 string 真实姓名
		* @param mobile 必选 string 电话号码
		* @return {"code":0,"data":{"id":2,"user_name":"jaxchan2","nickname":"","status":"","avatar":"default.png","created_at":"2020-06-12 19:38"},"msg":""}
		* @remark 无备注
		 */
		v1.POST("user/register", api.UserRegister)

		/**
		* showdoc
		* @catalog 用户模块
		* @title 抽奖用户登录
		* @description 抽奖用户的注册api
		* @method post
		* @url /api/v1/user/login
		* @param user_name 必选 string 用户名
		* @param password 必选 string 密码
		* @return {"code":0,"data":{"id":1,"user_name":"jaxchan","avatar":"default.png","created_at":"2020-06-12 19:03"},"msg":""}
		* @remark 无备注
		 */
		v1.POST("user/login", api.UserLogin)

		/*
			以下是关于奖品模块的api
		*/

		/**
		* showdoc
		* @catalog 奖品相关
		* @title 添加奖品
		* @description 添加奖品的api。
		* @method post
		* @url /api/v1/prize
		* @param activity_id 必选 int 活动id
		* @param level 必选 string 中奖等级
		* @param name 必选 string 奖品名字
		* @param prob 必选 double 中奖概率
		* @param all_num 必选 int 总个数
		* @param surplus_num 必选 int 剩余个数
		* @param image 可选 string 图片url地址
		* @param if_win 必选 int 判断本次是否中奖 1为不中奖 2为中奖
		* @return {"code":0,"data":{"id":3,"activity_id":2,"level":"0","name":"帅哥陈亮","prob":0.8,"all_num":10,"surplus_num":1,"image":"PrizeImage_default.png","if_win":1,"created_at":"2020-06-14 23:41"},"msg":""}
		 */
		v1.POST("prize", api.PrizePost)

		/**
		* showdoc
		* @catalog 奖品相关
		* @title 某个活动的奖品列表
		* @description 查看奖品列表的api。
		* @method get
		* @param activity_id 必选 int 活动主键
		* @url /api/v1/prizes/:id
		* @return {"code":0,"data":{"items":[{"id":1,"activity_id":1,"level":"一等奖","name":"葫芦娃1","prob":0.1,"all_num":100,"surplus_num":98,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":2,"activity_id":1,"level":"一等奖","name":"葫芦娃2","prob":0.1,"all_num":100,"surplus_num":99,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":3,"activity_id":1,"level":"一等奖","name":"葫芦娃3","prob":0.1,"all_num":100,"surplus_num":96,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":4,"activity_id":1,"level":"二等奖","name":"葫芦娃4","prob":0.1,"all_num":100,"surplus_num":97,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":5,"activity_id":1,"level":"二等奖","name":"葫芦娃5","prob":0.1,"all_num":100,"surplus_num":95,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":6,"activity_id":1,"level":"二等奖","name":"葫芦娃6","prob":0.1,"all_num":100,"surplus_num":99,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":7,"activity_id":1,"level":"三等奖","name":"葫芦娃7","prob":0.1,"all_num":100,"surplus_num":94,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":8,"activity_id":1,"level":"三等奖","name":"葫芦娃8","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":9,"activity_id":1,"level":"三等奖","name":"葫芦娃9","prob":0.1,"all_num":100,"surplus_num":99,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":10,"activity_id":1,"level":"不中奖","name":"葫芦娃10","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":1,"created_at":"2020-06-25 18:30"}],"total":10},"msg":""}
		 */
		v1.GET("prizes/:id", api.ListPrizes)

		/**
		* showdoc
		* @catalog 奖品相关
		* @title 某个活动的中奖记录
		* @description 查看某个活动的中奖列表的api。
		* @method post
		* @param activity_id 必选 int 活动主键
		* @param limit 可选 int 分页：每页大小
		* @param offset 可选 int 分页：开始值
		* @url /api/v1/activityprizes
		* @return {"code":0,"data":{"items":[{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:19","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:25","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:25","prize_name":"豪华别墅","prize_level":"一等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:26","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:19","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:25","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:25","prize_name":"豪华别墅","prize_level":"一等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"},{"user_name":"jaxchan","real_name":"陈亮","mobile":"15918897888","create_at":"2020-06-17 15:26","prize_name":"苹果5s","prize_level":"二等奖","province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10"}],"total":8},"msg":""}
		 */
		v1.POST("activityprizes", api.ListActivityPrizes)

		/*
			以下是活动相关的api
		*/

		/**
		* showdoc
		* @catalog 活动相关
		* @title 添加活动
		* @description 添加活动的api。
		* @method post
		* @url /api/v1/activity
		* @param game_manager_id 必选 int 管理员id
		* @param name 必选 string 活动名称
		* @param title 必选 string 活动title
		* @param start_time 必选 string
		* @param end_time 必选 string
		* @param limit_type 必选 int (1:每日抽奖次数限制,2:总抽奖次数限制)',
		* @param limit_num 必选 int '限制的抽奖次数',
		* @param rule_text 必选 string '活动规则介绍'，换行符使用:\r\n
		* @param banner_image 必选 string 头图
		* @param lottery_image 必选 string 抽奖图
		* @param background_color 必选 string 背景颜色
		* @param virtual_num 必选 int 虚拟参与者个数'
		* @return
		 */
		v1.POST("activity", api.ActivityPost)

		/**
		* showdoc
		* @catalog 活动相关
		* @title 活动列表
		* @description 查看活动列表的api。
		* @method post
		* @url /api/v1/activities
		* @param limit 可选 int 分页：每页大小
		* @param offset 可选 int 分页：开始值
		* @return {"code":0,"data":{"items":[{"id":1,"game_manager_id":2,"name":"社团抽奖活动","title":"1111title","type":0,"state":10,"mode":1,"limit_type":1,"limit_num":10,"rule_text":"只有美女才可以抽奖","foreground_image":"ForegroundImage_default.png","background_image":"BackgroundImage_default.png","virtual_num":888,"created_at":"2020-06-15 13:21"},{"id":2,"game_manager_id":2,"name":"社团抽奖活动","title":"1111title","type":0,"state":10,"mode":1,"limit_type":1,"limit_num":10,"rule_text":"只有美女才可以抽奖","foreground_image":"ForegroundImage_default.png","background_image":"BackgroundImage_default.png","virtual_num":888,"created_at":"2020-06-15 13:25"}],"total":2},"msg":""}
		 */
		v1.POST("activities", api.ListActivities)

		/**
		* showdoc
		* @catalog 活动相关
		* @title 活动详细页
		* @description 查看 活动详细的api。
		* @method get
		* @url /api/v1/activity/id
		* @param id 必选 int 主键
		* @return {"code":0,"data":{"id":1,"game_manager_id":1,"game_prizes":[{"id":1,"activity_id":1,"level":"一等奖","name":"葫芦娃1","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":2,"activity_id":1,"level":"一等奖","name":"葫芦娃2","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":3,"activity_id":1,"level":"一等奖","name":"葫芦娃3","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":4,"activity_id":1,"level":"二等奖","name":"葫芦娃4","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":5,"activity_id":1,"level":"二等奖","name":"葫芦娃5","prob":0.1,"all_num":100,"surplus_num":98,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":6,"activity_id":1,"level":"二等奖","name":"葫芦娃6","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":7,"activity_id":1,"level":"三等奖","name":"葫芦娃7","prob":0.1,"all_num":100,"surplus_num":98,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":8,"activity_id":1,"level":"三等奖","name":"葫芦娃8","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":9,"activity_id":1,"level":"三等奖","name":"葫芦娃9","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":10,"activity_id":1,"level":"不中奖","name":"葫芦娃10","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":1,"created_at":"2020-06-25 18:30"}],"name":"活动大转盘","title":"软件工程毕业晚会抽奖","start_time":"2020-06-27 15:06","end_time":"2020-07-24 15:06","limit_type":1,"limit_num":10,"rule_text":"1","banner_image":"https://cvte-oss.oss-cn-shenzhen.aliyuncs.com/banner_default.png","lottery_image":"https://cvte-oss.oss-cn-shenzhen.aliyuncs.com/lottery_default.png","background_color":"#FF6412","virtual_num":10,"created_at":"2020-06-24 13:43","participate_num":7},"msg":""}
		 */
		v1.GET("activity/:id", api.ShowActivity)

		/**
		* showdoc
		* @catalog 活动相关
		* @title 修改活动
		* @description 修改活动的api。
		* @method post
		* @url /api/v1/activity/id
		* @param id 必选 int 主键
		* @param title 可选 string 活动标题
		* @param limit_num 可选 int 限制的抽奖次数
		* @param rule_text 可选 string 活动规则介绍
		* @param start_time 可选 string
		* @param end_time 可选 string
		* @return {"code":0,"data":{"id":1,"game_manager_id":1,"game_prizes":[{"id":1,"activity_id":1,"level":"一等奖","name":"葫芦娃1","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":2,"activity_id":1,"level":"一等奖","name":"葫芦娃2","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":3,"activity_id":1,"level":"一等奖","name":"葫芦娃3","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":4,"activity_id":1,"level":"二等奖","name":"葫芦娃4","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":5,"activity_id":1,"level":"二等奖","name":"葫芦娃5","prob":0.1,"all_num":100,"surplus_num":98,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":6,"activity_id":1,"level":"二等奖","name":"葫芦娃6","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":7,"activity_id":1,"level":"三等奖","name":"葫芦娃7","prob":0.1,"all_num":100,"surplus_num":98,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":8,"activity_id":1,"level":"三等奖","name":"葫芦娃8","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":9,"activity_id":1,"level":"三等奖","name":"葫芦娃9","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":2,"created_at":"2020-06-25 18:30"},{"id":10,"activity_id":1,"level":"不中奖","name":"葫芦娃10","prob":0.1,"all_num":100,"surplus_num":100,"image":"https://qdtalk.com/wp-content/uploads/2017/09/1-2.png","if_win":1,"created_at":"2020-06-25 18:30"}],"name":"活动大转盘","title":"软件工程毕业晚会抽奖","start_time":"2020-06-27 15:06","end_time":"2020-07-24 15:06","limit_type":1,"limit_num":10,"rule_text":"1","banner_image":"https://cvte-oss.oss-cn-shenzhen.aliyuncs.com/banner_default.png","lottery_image":"https://cvte-oss.oss-cn-shenzhen.aliyuncs.com/lottery_default.png","background_color":"#FF6412","virtual_num":10,"created_at":"2020-06-24 13:43","participate_num":7},"msg":""}
		 */
		v1.POST("activity/:id", api.UpdateActivity)

		/*
			以下是PC业务相关的api
		*/
		/**
		* showdoc
		* @catalog PCWeb业务相关
		* @title 获取传播数据
		* @description 获取传播数据的api。
		* @method post
		* @url /api/v1/graphdata
		* @param id 必选 int 活动主键
		* @param start_time 必选 string 表格起始时间
		* @param end_time 必选 string 表格终止时间
		* @return {"code":0,"data":[{"period":"2020-06-27 10:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 11:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 12:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 13:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 14:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 15:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 16:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}},{"period":"2020-06-27 17:00","uint_data":{"view_count":9,"participate_count":6,"win_count":6}},{"period":"2020-06-27 18:00","uint_data":{"view_count":17,"participate_count":12,"win_count":12}},{"period":"2020-06-27 19:00","uint_data":{"view_count":0,"participate_count":0,"win_count":0}}],"msg":""}
		 */
		v1.POST("graphdata", api.GetGraphData)

		/*
			以下是关于抽奖模块，需要登录
		*/
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			/*
				以下是MOBILE前端业务相关的api
			*/

			/**
			* showdoc
			* @catalog MobileWeb业务相关
			* @title 进行抽奖
			* @description 某用户在某活动的进行抽奖的api。
			* @method post
			* @url /api/v1/drawlottery
			* @param game_user_id 必选 int 用户主键
			* @param activity_id 必选 int 活动主键
			* @return {"code":0,"data":{"id":7,"activity_id":1,"level":"不中奖","name":"谢谢惠顾","prob":0.6,"all_num":1000,"surplus_num":1000,"image":"PrizeImage_default.png","if_win":1,"created_at":"2020-06-17 14:33"},"msg":""}
			* @remark 活动必须处于运行期间，才可以进行抽奖。否则返回相应错误。
			 */
			auth.POST("drawlottery", api.DrawLottery)

			/**
			* showdoc
			* @catalog MobileWeb业务相关
			* @title 剩余抽奖次数
			* @description 某用户在某活动的剩余抽奖次数的api。
			* @method post
			* @url /api/v1/surplustimes
			* @param game_user_id 必选 int 用户主键
			* @param activity_id 必选 int 活动主键
			* @return {"code":0,"data":3,"msg":""}
			 */
			auth.POST("surplustimes", api.GetSurplusTime)

			/*
				以下是关于地址模块的api
			*/

			/**
			* showdoc
			* @catalog 地址相关
			* @title 添加地址
			* @description 添加地址的api。
			* @method post
			* @url /api/v1/address
			* @param game_user_id 必选 int 外码
			* @param province 必选 string 省份
			* @param real_name 必选 string 真实姓名
			* @param mobile 必选 string 手机号码
			* @param city 必选 string 市
			* @param district 必选 string 区
			* @param detail 必选 string 详细地址
			* @return {"code":0,"data":{"id":4,"game_user_id":1,"province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10","created_at":"2020-06-16 14:54"},"msg":""}
			 */
			auth.POST("address", api.AddressPost)

			/**
			* showdoc
			* @catalog 地址相关
			* @title 查看某个用户的地址
			* @description 查看某个用户的地址api。
			* @method get
			* @url /api/v1/address/id
			* @param id 必选 int 用户主键
			* @return {"code":0,"data":{"id":4,"game_user_id":1,"province":"广东省","city":"广州市","district":"番禺区","detail":"华南理工大学C10","created_at":"2020-06-16 14:54"},"msg":""}
			 */
			auth.GET("address/:id", api.ShowAddress)

			/**
			* showdoc
			* @catalog 奖品相关
			* @title 某个用户的中奖记录
			* @description 查看某个用户在某个活动的中奖记录的api。
			* @method post
			* @param game_user_id 必选 int 用户主键
			* @param activity_id 必选 int 活动主键
			* @url /api/v1/userprizes
			* @return {"code":0,"data":{"items":[{"id":8,"activity_id":1,"level":"二等奖","name":"苹果5s","prob":0.3,"all_num":10,"surplus_num":4,"image":"PrizeImage_default.png","if_win":2,"created_at":"2020-06-17 14:33"},{"id":9,"activity_id":1,"level":"一等奖","name":"豪华别墅","prob":0.1,"all_num":2,"surplus_num":1,"image":"PrizeImage_default.png","if_win":2,"created_at":"2020-06-17 14:33"}],"total":2},"msg":""}
			 */
			auth.POST("userprizes", api.ListUserPrizes)

		}
	}
	return r
}
