<script lang="ts">
    import QuizCard from "../../lib/QuizCard.svelte";
    import Button from "../../lib/Button.svelte";

    import type { Quiz } from "../../model/quiz";
    import { apiService } from "../../service/api";
    import { push } from 'svelte-spa-router'; // if not already imported

  
    let quizzes: Quiz[] = [];

    (async function () {
        quizzes = await apiService.getQuizzes();
    })();

      function createQuiz():void {
        push("/create");
    }

</script>

<div class="p-8">
    <h2 class="text-4xl font-bold">Your quizzes</h2>

    
    <div class="flex flex-col gap-2 mt-4">
        {#each quizzes as quiz (quiz.id)}
        <QuizCard on:host {quiz} />
        {/each}
    </div>
    <div class="flex flex-col gap-2 mt-4">
        <Button on:click={createQuiz}>Create New Quiz</Button>
    </div>
</div>
