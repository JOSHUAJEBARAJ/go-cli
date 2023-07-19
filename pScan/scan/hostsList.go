package scan

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var (
	ErrExists    = errors.New("Host already in the list")
	ErrNotExists = errors.New("Host not in the list")
)

type HostList struct {
	Hosts []string
}

func (h1 *HostList) search(host string) (bool, int) {
	sort.Strings(h1.Hosts)
	i := sort.SearchStrings(h1.Hosts, host)
	if i < len(h1.Hosts) && h1.Hosts[i] == host {
		return true, i
	}

	return false, -1

}

func (h1 *HostList) Add(host string) error {

	if found, _ := h1.search(host); found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}
	h1.Hosts = append(h1.Hosts, host)
	return nil
}

func (h1 *HostList) Remove(host string) error {
	if found, i := h1.search(host); found {
		h1.Hosts = append(h1.Hosts, h1.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("%w: %s", ErrNotExists, host)
}

// function to load the hostfile

func (h1 *HostList) Load(hostFile string) error {

	f, err := os.Open(hostFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		h1.Hosts = append(h1.Hosts, scanner.Text())
	}
	return nil
}

func (h1 *HostList) Save(hostFile string) error {
	output := ""
	for _, h := range h1.Hosts {
		output += fmt.Sprintln(h)
	}
	return ioutil.WriteFile(hostFile, []byte(output), 0644)
}
