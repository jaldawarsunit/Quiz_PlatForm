<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import { apiService } from "../../service/api";
    import { push } from "svelte-spa-router";

    let quizName: string = "";

    async function CreateQuiz() {
        if (quizName.trim() === "") return;
        console.log("Creating quiz with name:", quizName);
        // Call API service with name and empty questions array
        const created = await apiService.createQuiz({ name: quizName, questions: [] });
        console.log("Created quiz:", created);
        push("/host");
        // if (created && created.id) {
            // push(`/edit/${created}`);
        // } else {
            // Optionally handle error here
            // console.error("Failed to create quiz or missing id.");
        // }
  
    }

    function goBack() {
        push("/host");
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-purple-100">
    <div class="bg-white p-6 rounded shadow-md flex flex-col gap-4 w-full max-w-md">
        <h2 class="text-2xl font-bold text-center">Create New Quiz</h2>

        <input
            class="border border-gray-300 p-2 rounded"
            placeholder="Enter quiz name"
            bind:value={quizName}
        />

        <div class="flex justify-between">
            <Button on:click={CreateQuiz}>Create</Button>
            <div class="bg-gray-300 text-black">
                <Button on:click={goBack}>Back</Button>
            </div>
        </div>
    </div>
</div>
