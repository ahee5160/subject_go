package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:ahee5160@tcp(10.227.28.222:3306)/ahee?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
}

func SelectStorageCluster(clusterID int) (*StorageCluster, error) {
	sc := &StorageCluster{}
	rows := DB.QueryRow("select cluster, galaxy_id, psm_name from storage_cluster where id = ?", clusterID)
	err := rows.Scan(&sc.Cluster, &sc.GalaxyID, &sc.PSMName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, fmt.Sprintf("select storage_cluster failed, cluster id %d", clusterID))
	}
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("select error, cluster id %d", clusterID))
	}
	return sc, nil
}

type StorageCluster struct {
	Cluster  string `json:"cluster"`
	GalaxyID int    `json:"galaxy_id"`
	PSMName  string `json:"psm_name"`
}

func main() {
	// restful operation
	op := "get|post|put|delete"
	InitDB()
	sc, err := SelectStorageCluster(1)
	if errors.Is(err, sql.ErrNoRows) {
		if op == "get" {
			// 查询成功，但是数据数据为空
		} else if op == "put" || op == "delete" {
			// 操作失败，记录日志
		} else {
			// 插入成功
		}

	}
	if err != nil {
		fmt.Printf("original error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: %+v", err)
	} else {
		fmt.Printf("storage cluster: %+v", sc)
	}
}
