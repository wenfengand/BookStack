package models

import (
	"github.com/TruthHun/BookStack/conf"
	"github.com/astaxie/beego/orm"
)

// Option struct .
type Option struct {
	OptionId    int    `orm:"column(option_id);pk;auto;unique;" json:"option_id"`
	OptionTitle string `orm:"column(option_title);size(500)" json:"option_title"`
	OptionName  string `orm:"column(option_name);unique;size(80)" json:"option_name"`
	OptionValue string `orm:"column(option_value);type(text);null" json:"option_value"`
	Remark      string `orm:"column(remark);type(text);null" json:"remark"`
}

// TableName 获取对应数据库表名.
func (m *Option) TableName() string {
	return "options"
}

// TableEngine 获取数据使用的引擎.
func (m *Option) TableEngine() string {
	return "INNODB"
}

func (m *Option) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewOption() *Option {
	return &Option{}
}

func (p *Option) Find(id int) (*Option, error) {
	o := orm.NewOrm()

	p.OptionId = id

	if err := o.Read(p); err != nil {
		return p, err
	}
	return p, nil
}

func (p *Option) FindByKey(key string) (*Option, error) {
	o := orm.NewOrm()
	if err := o.QueryTable(p).Filter("option_name", key).One(p); err != nil {
		return p, err
	}
	return p, nil
}

func GetOptionValue(key, def string) string {

	if option, err := NewOption().FindByKey(key); err == nil {
		return option.OptionValue
	}
	return def
}

func (p *Option) InsertOrUpdate() error {

	o := orm.NewOrm()

	var err error

	if p.OptionId > 0 || o.QueryTable(p.TableNameWithPrefix()).Filter("option_name", p.OptionName).Exist() {
		_, err = o.Update(p)
	} else {
		_, err = o.Insert(p)
	}
	return err
}

func (p *Option) InsertMulti(option ...Option) error {

	o := orm.NewOrm()

	_, err := o.InsertMulti(len(option), option)
	return err
}

func (p *Option) All() ([]*Option, error) {
	o := orm.NewOrm()
	var options []*Option

	_, err := o.QueryTable(p.TableNameWithPrefix()).All(&options)

	if err != nil {
		return options, err
	}
	return options, nil
}

func (m *Option) Init() error {

	o := orm.NewOrm()

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ENABLED_REGISTER").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "ENABLED_REGISTER"
		option.OptionTitle = "是否启用注册"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ENABLE_DOCUMENT_HISTORY").Exist() {
		option := NewOption()
		option.OptionValue = "100"
		option.OptionName = "ENABLE_DOCUMENT_HISTORY"
		option.OptionTitle = "版本控制"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ENABLED_CAPTCHA").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "ENABLED_CAPTCHA"
		option.OptionTitle = "是否启用验证码"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ENABLE_ANONYMOUS").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "ENABLE_ANONYMOUS"
		option.OptionTitle = "启用匿名访问"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "SITE_NAME").Exist() {
		option := NewOption()
		option.OptionValue = "BookStack"
		option.OptionName = "SITE_NAME"
		option.OptionTitle = "站点名称"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ICP").Exist() {
		option := NewOption()
		option.OptionValue = ""
		option.OptionName = "ICP"
		option.OptionTitle = "网站备案"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "TONGJI").Exist() {
		option := NewOption()
		option.OptionValue = ""
		option.OptionName = "TONGJI"
		option.OptionTitle = "站点统计"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "SPIDER").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "SPIDER"
		option.OptionTitle = "采集器，是否只对管理员开放"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ELASTICSEARCH_ON").Exist() {
		option := NewOption()
		option.OptionValue = "false"
		option.OptionName = "ELASTICSEARCH_ON"
		option.OptionTitle = "是否开启全文搜索"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "ELASTICSEARCH_HOST").Exist() {
		option := NewOption()
		option.OptionValue = "http://localhost:9200/"
		option.OptionName = "ELASTICSEARCH_HOST"
		option.OptionTitle = "ElasticSearch Host"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "DEFAULT_SEARCH").Exist() {
		option := NewOption()
		option.OptionValue = "book"
		option.OptionName = "DEFAULT_SEARCH"
		option.OptionTitle = "默认搜索"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "SEARCH_ACCURACY").Exist() {
		option := NewOption()
		option.OptionValue = "50"
		option.OptionName = "SEARCH_ACCURACY"
		option.OptionTitle = "搜索精度"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "LOGIN_QQ").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "LOGIN_QQ"
		option.OptionTitle = "是否允许使用QQ登录"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "LOGIN_GITHUB").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "LOGIN_GITHUB"
		option.OptionTitle = "是否允许使用Github登录"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "LOGIN_GITEE").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "LOGIN_GITEE"
		option.OptionTitle = "是否允许使用码云登录"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	// 书籍公开发布，是否需要执行审核
	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "CHECK_BOOK").Exist() {
		option := NewOption()
		option.OptionValue = "true"
		option.OptionName = "CHECK_BOOK"
		option.OptionTitle = "是否开启书籍审核"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	if !o.QueryTable(m.TableNameWithPrefix()).Filter("option_name", "RELATE_BOOK").Exist() {
		option := NewOption()
		option.OptionValue = "0"
		option.OptionName = "RELATE_BOOK"
		option.OptionTitle = "是否开始关联书籍"
		if _, err := o.Insert(option); err != nil {
			return err
		}
	}

	return nil
}
