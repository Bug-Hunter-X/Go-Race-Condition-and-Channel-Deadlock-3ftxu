# Go Race Condition and Channel Deadlock Example

This repository contains a simple Go program that illustrates a potential race condition and subsequent deadlock involving channels and wait groups.

## Bug Description
The program launches multiple goroutines, each sending a value to a channel. The main goroutine receives these values.  However, the `wg.Wait()` and channel operations are not properly coordinated, leading to a potential deadlock if the channel is closed too early or the main goroutine attempts to read from the channel after all senders have finished.