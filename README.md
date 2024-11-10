# JG.Go.Quiz.Game
JG.Go.Quiz.Game
It is a small example of quiz in GO + Json (external source data) with the theme **Super Heroes**

Super heroes game quiz for kids that is learning GO + Command line knowledge

<h2>Steps to start the game:</h2>

<ul>
  <li>Execute the command in the terminal from the root folder (JG.quiz): <b>go run main.go</b></li>
  <li>Inform how many questions you want to answer in this quiz</li>
  <li>Inform your name and email</li>
  <li>Start the quiz</li>
</ul> 

![image](https://github.com/user-attachments/assets/e3b2cef4-f71c-44d0-a79d-e8e319e9d900)

<ul>
<li>See the results at the end
</li>
</ul>

![image](https://github.com/user-attachments/assets/ac41e5e1-7a8f-4258-aa96-5e4eab967d9f)

<h3>Update the data.json file and incrementing with more questions</h3>

To close the application, use the command CTRL + C in your keyboard from Windows or Command + C using Mac


<u>
<li>You can add extra questions updating the data.json file</li> 
  
```json
[
  {
    "Title": "Who is the superhero with a hammer that only he* can lift?",
    "Answers": [
      {"Title": "A - Hulk", "IsRight": false, "Code": "A"},
      {"Title": "B - Thor", "IsRight": true, "Code": "B"},
      {"Title": "C - Spider", "IsRight": false, "Code": "C"}
    ]
  },
  {
    "Title": "Which superhero has a suit that can fly and shoot lasers from his hands?",
    "Answers": [
      {"Title": "A - Wolverine", "IsRight": false, "Code": "A"},
      {"Title": "B - Black Panther", "IsRight": false, "Code": "B"},
      {"Title": "C - Iron Man", "IsRight": true, "Code": "C"}
    ]
  }
]
