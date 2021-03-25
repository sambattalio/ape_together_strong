# ape_together_strong
![monkey](https://chadpaste.com/f/wtcwkifwlo.jpg)

## Usage
To deploy this project, just download and unpack the repo and run the main Go script with `go run main.go`.

To use this project, simply log into the website and use the input boxes to design a workout plan, then hit the gym and follow it! To design a plan, for each exercise you want in your plan, simply enter the exercise you want to do, the weight you are planning on using, and the day you want to do it, then hit submit.

## Goal and Unique Features
The goal of this project was to create a website that people can use to design workouts in advance of going to the gym and to track their progress across workouts. We felt that many people refrain from going to the gym because they don't know what to do once they get there and because they don't feel like they are making progress. Our project allows users to design a workout plan from the comfort of their home so that they know exactly what to do once they're in the gym, and it tracks the weights they are using in order to display a graph of the user's progress to them, thereby addressing both of those issues. Of course, if a user is not making progress the graph won't help them much, but we hope that even for such users seeing a flat or declining graph will inspire them to redouble their efforts.

We chose to use Go because we felt that its unique features mean that it lends itself well to use as a web service. Goroutines, which are essentially easy-to-use, lightweight threads, are extremely well suited for dealing with REST requests. When a request is received, our artifact simply spins off a goroutine to handle it. This is very useful, because it allows the main thread to go back to receiving requests and serving responses. Additionally, many REST requests, both in general and in our specific case, do not require much computation time to service, meaning that the savings from using lightweight goroutines instead of traditional threads is significant, because thread overhead would otherwise account for a significant portion of our resource usage. This second benefit is particularly important for a web service, as we will likely need to handle requests from many users at once, meaning that we will likely need to have many goroutines (which would otherwise be threads) running at once.
