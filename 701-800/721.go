package _01_800

import "sort"

/**

 */
func accountsMerge(accounts [][]string) [][]string {
	l := len(accounts)

	// 构建初始并查集
	uf_set := make([]int, l)
	find := func(id int) int {
		for id != uf_set[id] {
			uf_set[id] = uf_set[uf_set[id]]
			id = uf_set[id]
		}
		return id
	}
	union := func(id1, id2 int) {
		set_id1, set_id2 := find(id1), find(id2)
		if set_id1 < set_id2 {
			uf_set[set_id2] = set_id1
		} else if set_id1 > set_id2 {
			uf_set[set_id1] = set_id2
		}
	}
	for id := range uf_set {
		uf_set[id] = id
	}

	// 对账户关系进行集合并查
	mail2ids := map[string][]int{}
	for id, account := range accounts {
		for _, mail := range account[1:] {
			mail2ids[mail] = append(mail2ids[mail], id)
		}
	}
	order_mails := []string{}
	for mail, ids := range mail2ids {
		order_mails= append(order_mails, mail)
		for _, id := range ids[1:] {
			union(id, ids[0])
		}
	}
	sort.Strings(order_mails)

	// 合并相同集合的账户
	mergeAccounts := make([][]string, l)
	for _, mail := range order_mails {
		id := find(mail2ids[mail][0])
		if len(mergeAccounts[id]) == 0 {
			name := accounts[id][0]
			mergeAccounts[id] = append(mergeAccounts[id], name)
		}
		mergeAccounts[id] = append(mergeAccounts[id], mail)
	}
	res := [][]string{}
	for _, account := range mergeAccounts {
		if len(account) != 0 {
			res = append(res, account)
		}
	}
	return res
}