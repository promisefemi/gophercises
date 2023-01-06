package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open("task.db", 0777, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		taskb := tx.Bucket(taskBucket)
		id64, _ := taskb.NextSequence()
		id = int(id64)
		key := itob(id)
		return taskb.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTask() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		taskb := tx.Bucket(taskBucket)
		series := taskb.Cursor()

		for key, value := series.First(); key != nil; key, value = series.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(key),
				Value: string(value),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		taskb := tx.Bucket(taskBucket)
		return taskb.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
