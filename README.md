# Rock✊ Paper🫱 Scissor✌️ Game in Go

Welcome to the **Rock✊ Paper🫱 Scissor✌️ Game**, a simple, fun, and interactive CLI-based implementation written in Go! Play against the computer in a best-of-3 match and test your luck and strategy.

---

## 🎮 **Game Rules**
- Rock beats Scissor.
- Scissor beats Paper.
- Paper beats Rock.
- If both players choose the same, the round ends in a draw.

---

## 🛠️ **How to Run**

### Prerequisites
- Go installed on your machine. Download it [here](https://golang.org/dl/).

### Steps
1. Clone the repository or copy the code into a local file named `main.go`.
2. Open your terminal and navigate to the directory containing the file.
3. Run the game using the following command:
   ```bash
   go run main.go
   ```

---

## 🕹️ **How to Play**
1. The game consists of **3 rounds**.
2. Each round:
   - You will be prompted to enter your choice: `Rock`, `Paper`, or `Scissor`.
   - The computer will randomly select its choice.
3. After each round:
   - The winner of the round is announced.
   - Scores are updated.

4. After 3 rounds, the final winner is declared based on the scores.

---

## 📝 **Sample Output**
```
Rock✊  Paper🫱  Scissor✌️  Game
🥏Round 1
Enter your choice (Rock, Paper, scissor): Rock
Computer choice: Paper🫱
Computer won the round.🛠️

🥏Round 2
Enter your choice (Rock, Paper, scissor): Scissor
Computer choice: Scissor✌️
It's a draw.🤝

🥏Round 3
Enter your choice (Rock, Paper, scissor): Paper
Computer choice: Rock✊
You won the round.✨

Game Over!
It's a draw match.😶‍🌫️
```

---

## 🤖 **Behind the Scenes**
- **Randomization:** The computer's choice is randomly generated using `math/rand`.
- **Case-Insensitive Input:** The player's input is converted to lowercase for comparison.
- **Winning Logic:** The game determines the winner using `switch` statements and conditions.

---

## 📂 **Project Structure**
```plaintext
📁 Rock-Paper-Scissor-Game
    └── main.go  # Core game logic
```

---

## 🤩 **Features**
- **Interactive Gameplay:** Fully command-line-based and user-friendly.
- **Best-of-3 Match:** Ensures a fair and engaging experience.
- **Fun Emojis:** Adds a touch of fun to the gameplay.

---

## 📜 **License**
This project is licensed under the MIT License. Feel free to use, modify, and distribute it as you wish.

---

## 🚀 **Future Enhancements**
- Add more rounds or allow the user to choose the number of rounds.
- Introduce a GUI for a more engaging experience.
- Support for multiplayer mode.

---

## 🧑‍💻 **Contributions**
Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

---

Enjoy the game and may the odds be in your favor! ✨
