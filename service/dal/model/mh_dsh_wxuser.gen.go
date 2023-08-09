// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/go-mumu/cs-go/library/common/datetime"

const TableNameWxuser = "mh_dsh_wxuser"

// Wxuser mapped from table <mh_dsh_wxuser>
type Wxuser struct {
	Userid            int32              `gorm:"column:userid;type:int(11);primaryKey;autoIncrement:true;comment:用户id" json:"userid"`                                            // 用户id
	Mid               uint64             `gorm:"column:mid;type:bigint(20) unsigned;not null;comment:main id 用户中心uid" json:"mid"`                                                // main id 用户中心uid
	Nickname          string             `gorm:"column:nickname;type:varchar(100);not null;comment:用户昵称" json:"nickname"`                                                        // 用户昵称
	Openid            string             `gorm:"column:openid;type:varchar(100);not null;comment:微信openid" json:"openid"`                                                        // 微信openid
	Unionid           string             `gorm:"column:unionid;type:char(50);not null;comment:unionid" json:"unionid"`                                                           // unionid
	CsOpenid          string             `gorm:"column:cs_openid;type:varchar(100);not null;comment:财商openid" json:"cs_openid"`                                                  // 财商openid
	XcxOpenid         string             `gorm:"column:xcx_openid;type:varchar(100);not null;comment:小程序openid" json:"xcx_openid"`                                               // 小程序openid
	Sex               int32              `gorm:"column:sex;type:int(11);not null;comment:默认0,1男,2女" json:"sex"`                                                                  // 默认0,1男,2女
	City              string             `gorm:"column:city;type:varchar(255);not null;comment:城市" json:"city"`                                                                  // 城市
	Province          string             `gorm:"column:province;type:varchar(255);not null;comment:省份" json:"province"`                                                          // 省份
	Points            int32              `gorm:"column:points;type:int(11);not null;comment:积分" json:"points"`                                                                   // 积分
	Long              int32              `gorm:"column:long;type:int(11);not null;comment:邀请人ID" json:"long"`                                                                    // 邀请人ID
	Longtime          *datetime.Datetime `gorm:"column:longtime;type:datetime;not null;default:0000-01-01 00:00:00;comment:邀请时间" json:"longtime"`                                // 邀请时间
	URL               string             `gorm:"column:url;type:varchar(255);not null;comment:头像" json:"url"`                                                                    // 头像
	Vip7              int32              `gorm:"column:vip7;type:tinyint(4);not null;comment:是否领取过vip" json:"vip7"`                                                              // 是否领取过vip
	Viptime           *datetime.Datetime `gorm:"column:viptime;type:datetime;not null;default:0000-01-01 00:00:00;comment:成为vip时间" json:"viptime"`                               // 成为vip时间
	Vipvalidity       *datetime.Datetime `gorm:"column:vipvalidity;type:datetime;not null;default:0000-01-01 00:00:00;comment:vip有效时间" json:"vipvalidity"`                       // vip有效时间
	Createtime        *datetime.Datetime `gorm:"column:createtime;type:datetime;not null;default:0000-01-01 00:00:00;comment:添加时间" json:"createtime"`                            // 添加时间
	SubChannel        *string            `gorm:"column:sub_channel;type:varchar(11);not null;default:1000;comment:关注渠道号 默认：1000" json:"sub_channel"`                             // 关注渠道号 默认：1000
	SubChannelTime    *datetime.Datetime `gorm:"column:sub_channel_time;type:datetime;not null;default:0000-01-01 00:00:00;comment:关注渠道修改时间" json:"sub_channel_time"`            // 关注渠道修改时间
	XcxSubChannel     *string            `gorm:"column:xcx_sub_channel;type:varchar(11);not null;default:1000;comment:小程序关注渠道号 默认：1000" json:"xcx_sub_channel"`                  // 小程序关注渠道号 默认：1000
	XcxSubChannelTime *datetime.Datetime `gorm:"column:xcx_sub_channel_time;type:datetime;not null;default:0000-01-01 00:00:00;comment:小程序关注渠道修改时间" json:"xcx_sub_channel_time"` // 小程序关注渠道修改时间
	BookChannel       *string            `gorm:"column:book_channel;type:varchar(11);not null;default:1000;comment:新渠道号 默认：1000" json:"book_channel"`                            // 新渠道号 默认：1000
	UserChannel       int32              `gorm:"column:user_channel;type:int(11);not null;comment:1小白营" json:"user_channel"`                                                     // 1小白营
	UserType          int32              `gorm:"column:user_type;type:tinyint(4);not null;comment:1裂变新用户" json:"user_type"`                                                      // 1裂变新用户
	ActivityName      string             `gorm:"column:activity_name;type:varchar(30);not null;comment:来源活动名称" json:"activity_name"`                                             // 来源活动名称
}

// TableName Wxuser's table name
func (*Wxuser) TableName() string {
	return TableNameWxuser
}
