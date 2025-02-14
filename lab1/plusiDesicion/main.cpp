#include <iostream>
#include <vector>
#include <stack>
#include <cmath>

struct Square {
    int x, y, w;  
};

class Board {
public:
    Board(int size) {
        this->size = size;
        this->defaultSquareSize = size / 2;
        this->curSquareSize = this->defaultSquareSize;
        Init();
    }

    void PlaceSquare(int x, int y) {
        if (y + curSquareSize > size || y < 0 || x + curSquareSize > size || x < 0) {
            return;
        }

        countSquares++;
        for (int i = y; i < y + curSquareSize; ++i) {
            for (int j = x; j < x + curSquareSize; ++j) {
                matrix[i][j] = countSquares;
            }
        }

        bestSolution.push({ x, y, curSquareSize });  
        curSquareSize = defaultSquareSize;
    }

    bool Fill(int minAmount) {
        for (int y = size / 2; y < size; ++y) {
            for (int x = size / 2; x < size; ++x) {
                if (!matrix[y][x]) {
                    if (countSquares >= minAmount) {
                        return false;
                    }

                    while (!CheckEnoughPlace(x, y)) {
                        curSquareSize--;
                    }
                    PlaceSquare(x, y);
                }
            }
        }
        return true;
    }

    void Backtrace() {
        Square lastSquare = bestSolution.top();
        while (bestSolution.size() > 3 && lastSquare.w == 1) {  
            DeleteSquare();
            lastSquare = bestSolution.top();
        }

        if (bestSolution.size() == 3) {
            running = false;
            return;
        }

        int x = lastSquare.x, y = lastSquare.y;
        curSquareSize = lastSquare.w - 1;  
        DeleteSquare();
        PlaceSquare(x, y);
    }

    void PrintCoords() const {
        std::cout << bestSolution.size() << "\n";
        std::stack<Square> tempStack = bestSolution;

        while (!tempStack.empty()) {
            const Square& square = tempStack.top();
            std::cout << square.x + 1 << " " << square.y + 1 << " " << square.w << "\n";  
            tempStack.pop();
        }
    }

    bool getRunning() const { 
        return running; 
    }

    int getSquaresAmount() const { 
        return countSquares; 
    }

    void setCurSquareSize(int value) { 
        if (value > 0) {
            curSquareSize = value; 
        }
    }

private:
    int size, defaultSquareSize, countSquares = 0, curSquareSize = 0;
    bool running = true;

    std::vector<std::vector<int>> matrix;
    std::stack<Square> bestSolution;

    void Init() {
        matrix.resize(size, std::vector<int>(size, 0));
    }

    bool CheckEnoughPlace(int startX, int startY) {
        if (startY + curSquareSize > size || startX + curSquareSize > size) {
            return false;
        }

        for (int y = startY; y < startY + curSquareSize; y++) {
            for (int x = startX; x < startX + curSquareSize; x++) {
                if (matrix[y][x]) {
                    return false;
                }
            }
        }

        return true;
    }

    void DeleteSquare() {
        Square lastSquare = bestSolution.top();
        bestSolution.pop();
        countSquares--;

        for (int y = lastSquare.y; y < lastSquare.y + lastSquare.w; y++) {  
            for (int x = lastSquare.x; x < lastSquare.x + lastSquare.w; x++) {  
                matrix[y][x] = 0;
            }
        }
    }
};

int main() {
    int n;
    std::cin >> n;

    if(n % 2 == 0) {
        std::cout << 4 << "\n";
        std::cout << 1 << " " << 1 << " " << n / 2 << "\n";
        std::cout << 1 + n / 2 << " " << 1 << " " << n / 2 << "\n";
        std::cout << 1 << " " << 1 + n / 2 << " " << n / 2 << "\n";
        std::cout << 1 + n / 2 << " " << 1 + n / 2 << " " << n / 2 << "\n";
        return 0;
    } else if(n % 3 == 0) {
        std::cout << 6 << "\n";

        std::cout << 1 << " " << 1 << " " << 2 * n / 3 << "\n";
        std::cout << 1 + 2 * n / 3 << " " << 1 << " " << n / 3 << "\n";
        std::cout << 1 << " " << 1 + 2 * n / 3 << " " << n / 3 << "\n";
        std::cout << 1 + 2 * n / 3 << " " << 1 + n / 3 << " " << n / 3 << "\n";
        std::cout << 1 + n / 3 << " " << 1 + 2 * n / 3 << " " << n / 3 << "\n";
        std::cout << 1 + 2 * n / 3 << " " << 1 + 2 * n / 3 << " " << n / 3 << "\n";
        return 0;
    } else if(n % 5 == 0) {
        std::cout << 8 << "\n";
        std::cout << 1 << " " << 1 << " " << 3 * n / 5 << "\n";
        std::cout << 1 + 3 * n / 5 << " " << 1 << " " << 2 * n / 5 << "\n";
        std::cout << 1 << " " << 1 + 3 * n / 5 << " " << 2 * n / 5 << "\n";
        std::cout << 1 + 3 * n / 5 << " " << 1 + 3 * n / 5 << " " << 2 * n / 5 << "\n";
        std::cout << 1 + 2 * n / 5 << " " << 1 + 3 * n / 5 << " " << n / 5 << "\n";
        std::cout << 1 + 2 * n / 5 << " " << 1 + 4 * n / 5 << " " << n / 5 << "\n";
        std::cout << 1 + 3 * n / 5 << " " << 1 + 2 * n / 5 << " " << n / 5 << "\n";
        std::cout << 1 + 4 * n / 5 << " " << 1 + 2 * n / 5 << " " << n / 5 << "\n";
        return 0;
    }

    Board board(n), minField(n);
    board.setCurSquareSize(std::ceil(double(n) / 2));
    board.PlaceSquare(0, 0);
    board.setCurSquareSize(n / 2);
    board.PlaceSquare(std::ceil(double(n) / 2), 0);
    board.PlaceSquare(0, std::ceil(double(n) / 2));

    int minAmount = n * n;

    while (board.getRunning()) {
        bool filledSuccessfully = board.Fill(minAmount);

        if (filledSuccessfully && board.getSquaresAmount() < minAmount) {
            minField = board;
            minAmount = board.getSquaresAmount();
        }
        
        board.Backtrace();
    }

    minField.PrintCoords();
    return 0;
}