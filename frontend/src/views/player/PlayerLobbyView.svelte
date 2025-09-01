<script lang="ts">

    import { createEventDispatcher } from "svelte";
    import Button from "../../lib/Button.svelte";
    import type { PlayerGame } from "../../service/player/player";
     function returnToLobby() {
        window.location.reload();
    }

</script>

<div class="bg-purple-400 min-h-screen w-full flex items-center justify-center">
    <div>
        <h2 class="text-white font-bold text-5xl text-center">Quiz</h2>
        <div class="flex flex-col gap-2 mt-10 items-center">
            Welcome to the game!
Do you see your name on the screen?
        </div>
        <div class="flex gap-4 mt-8 justify-center">
            <button class="bg-white px-6 py-2 rounded-full text-purple-600 font-semibold" on:click={returnToLobby}>
                Leave This Quiz
            </button>      
        </div>

    </div>
</div>
