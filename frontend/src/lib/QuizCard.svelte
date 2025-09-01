<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import Button from "./Button.svelte";
    import type { Quiz } from "../model/quiz";
    import { push } from "svelte-spa-router";

    const dispatch = createEventDispatcher();

    export let quiz: Quiz;

    function host(){
        dispatch("host", quiz);
    }
    
    function edit() {
        push(`/edit/${quiz.id}`);
    }
    async function del(){
        alert(`Are you sure you want to delete the quiz: ${quiz.name}?`);
        console.log("Delete quiz with id:", quiz.id);
        const response = await fetch(`http://localhost:3000/api/quiz/${quiz.id}`, {
            method: "DELETE",
        });
        // alert("response", response);

        if (response.ok) {
            alert("Quiz deleted successfully!");
            // redirect or refresh
            push("/host"); // if you want to go back to home or quizzes list
        } else {
            alert("Failed to delete quiz.");
        }
        //implemetn the lgic of delete that quzei
    }
</script>

<div class="flex justify-between items-center bg-white border p-4 rounded-xl">
    <p>{quiz.name}</p>
    <div class="flex gap-2 items-center">
        <div class="bg-white-500 text-white rounded px-1 py-1 hover:bg-green-400 transition">
            <Button on:click={host}>Host</Button>
        </div>
        <div class="bg-white-500 text-white rounded px-1 py-1 hover:bg-green-400 transition">
            <Button on:click={edit}>Edit</Button>
        </div>
        <div class="bg-white-500 text-white rounded px-1 py-1 hover:bg-red-400 transition">
            <Button on:click={del}>Delete</Button>

        </div>
    </div>
</div>