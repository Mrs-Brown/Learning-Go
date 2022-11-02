//Introducing mutex
package main

import "sync"  //pkg to use mutex

var mu sync.Mutex  //Declare mutex

func main() {
  mu.Lock()  //Lock the mutex
  defer mu.Unlock //Unlock mutex before returning
  //The lock is held until we return from the function
}
