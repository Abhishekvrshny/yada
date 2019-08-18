package downloader

import (
	"fmt"
	"math"
	"time"

	"github.com/Abhishekvrshny/yada/constants"
	"github.com/Abhishekvrshny/yada/utils"
)

type Concurrent struct {
	BaseDownloader
	jobChan    chan job
	resultChan chan bool
}

type job struct {
	filePath string
	url      string
}

func (c *Concurrent) Download() error {
	c.startTime = time.Now()
	err := utils.CreateDir(constants.DOWNLOAD_PATH + c.id)
	if err != nil {
		return err
	}
	for _, file := range c.urls {
		fileID, err := utils.UUID()
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%s%s/%s", constants.DOWNLOAD_PATH, c.id, fileID)

		c.jobChan <- job{filePath: filePath, url: file}
	}
	go c.waitForCompletion()
	return nil
}

func NewConcurrent(urls []string) (Downloader, error) {
	uuid, err := utils.UUID()
	if err != nil {
		return nil, err
	}
	c := &Concurrent{
		BaseDownloader{
			id:           uuid,
			urls:         urls,
			fileMap:      make(map[string]string),
			downloadType: constants.CONCURRENT,
		},
		make(chan job),
		make(chan bool),
	}

	numWorkers := math.Max(float64(len(urls)), constants.MAX_WORKERS_PER_DOWNLOAD)
	for i := 0; i < int(numWorkers); i++ {
		go c.spawnWorkers()
	}
	return c, nil
}

func (c *Concurrent) spawnWorkers() {
	for job := range c.jobChan {
		err := utils.DownloadFile(job.filePath, job.url)
		if err != nil {
			c.resultChan <- false
		} else {
			c.resultChan <- true
		}
	}
}

func (c *Concurrent) waitForCompletion() {
	status := constants.SUCCESSFUL
	counter := 0
	for result := range c.resultChan {
		if result == false {
			status = constants.FAILED
		}
		counter += 1
		if counter == len(c.urls) {
			close(c.jobChan)
			close(c.resultChan)
		}
	}
	c.status = status
	c.endTime = time.Now()
}
