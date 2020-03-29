package visitor_persistence

type VPersistence interface {
	Add(string, int64) error                              //增加一条访问记录
	List(string, []int64, int64) (map[int64]int64, error) //列出指定时间段的访问次数
}
