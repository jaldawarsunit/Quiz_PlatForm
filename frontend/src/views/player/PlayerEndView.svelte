<script lang="ts">
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { leaderboard } from "../../service/host/host";
    import type { PlayerGame } from "../../service/player/player";
    // import type { Player } from "../../model";

        function returnToLobby() {
        window.location.reload();
    }

</script>

<div class="flex justify-center bg-purple-500 min-h-screen w-full">
    <div class="mt-32">
        <h2 class="text-center text-white text-5xl font-bold">Game ended!</h2>
        <h2 class="text-center text-white text-3xl font-bold">Result Are On Host Screen</h2>
        <!-- <h2>{$player}</h2> -->
        <div class="flex gap-4 mt-8 justify-center">
            <button class="bg-white px-6 py-2 rounded-full text-purple-600 font-semibold" on:click={returnToLobby}>
                Return to Lobby
            </button>      
        </div>

         <br>
         <br>
        <h3 class="text-center text-white text-2xl mt-6">Thank you for playing!</h3>
        <!-- Output player name if available -->
        <!-- Example: -->
        <!-- <h4 class="text-center text-white text-xl mt-2">{$Player?.name}</h4> -->
        <!-- <div class="flex flex-wrap gap-2 mt-10"> -->
            <!-- <h3>Result Are On Host Screen</h3> -->
        <!-- </div> -->
    </div>
</div>
