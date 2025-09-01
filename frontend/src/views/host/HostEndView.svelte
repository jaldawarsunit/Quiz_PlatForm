<script lang="ts">


    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { leaderboard } from "../../service/host/host";



    function returnToLobby() {
        window.location.reload();
    }

    function downloadCSV() {
        const rows = [["Name", "Points"]];
        $leaderboard.forEach(player => {
            rows.push([player.name, player.points.toString()]);
        });

        const csvContent = rows.map(e => e.join(",")).join("\n");
        const blob = new Blob([csvContent], { type: "text/csv" });
        const url = URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = "quiz_results.csv";
        a.click();
        URL.revokeObjectURL(url);
    }
</script>

<div class="flex justify-center bg-purple-500 min-h-screen w-full">
    <div class="mt-32">
        <h2 class="text-center text-white text-5xl font-bold">Game ended!</h2>
        <div class="flex flex-wrap gap-2 mt-10">
            <Leaderboard finish={true} leaderboard={$leaderboard} />
        </div>

        <div class="flex gap-4 mt-8 justify-center">
            <button class="bg-white px-6 py-2 rounded-full text-purple-600 font-semibold" on:click={returnToLobby}>
                Return to Lobby
            </button>
         
                <button class="bg-white px-6 py-2 rounded-full text-green-600 font-semibold" on:click={downloadCSV}>
                    Download Results
                </button>
          
        </div>
    </div>
</div>
