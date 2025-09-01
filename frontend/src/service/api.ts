import type { Quiz } from "../model/quiz";

export class ApiService {
    async getQuizById(id: string): Promise<Quiz | null> {
        let response = await fetch(`http://localhost:3000/api/quizzes/${id}`);
        if (!response.ok) return null;
        return await response.json();
    }

    async getQuizzes(): Promise<Quiz[]> {
        let response = await fetch("http://localhost:3000/api/quizzes");
        if (!response.ok) {
            alert("Failed to fetch quizzes!");
            return [];
        }
        return await response.json();
    }

    async saveQuiz(quizId: string, quiz: Quiz) {
        const response = await fetch(`http://localhost:3000/api/quizzes/${quizId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(quiz)
        });

        if (!response.ok) {
            alert("Failed to save quiz!");
        }
    }

    async createQuiz(quiz: { name: string; questions: any[] }): Promise<Quiz | null> {
        const response = await fetch("http://localhost:3000/api/quiz", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(quiz)
        });

        if (!response.ok) {
            // console.log("api");
            // alert("api Failed to create quiz!");
            return null;
        }

        return await response.json();
    }
}

export const apiService = new ApiService();
