package main

import "context"

func main() {
	rootCtx := context.Background()
	childCtx1 := context.WithValue(rootCtx, "request_Id", "seekload")
	childCtx2, cancelFunc := context.WithCancel(childCtx1)
	childCtx3 := context.WithValue(rootCtx, "user_Id", "user_100")
	childCtx4 := context.WithValue(childCtx1, "token", "token_some")
}
/*


                    ┌────────┐
                    │   root │
               ┌────┴────────┴─────┐
               │                   │
               │                   │
               │                   │
           ┌───▼───┐          ┌────▼───┐
           │child1 │          │child3  │
       ┌───┴───────┴────┐     └────────┘
       │                │
       │                │
   ┌───▼───┐        ┌───▼───┐
   │child2 │        │child4 │
   └───────┘        └───────┘






*/