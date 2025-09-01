<script lang="ts">
  import Router from "svelte-spa-router";
  import HostView from "./views/host/HostView.svelte";
  import PlayerView from "./views/player/PlayerView.svelte";
  import EditQuizView from "./views/edit/EditQuizView.svelte";
  import CreateQuiz from "./views/create/CreateQuiz.svelte";

  let routes = {
    "/": PlayerView,
    "/host": HostView,
    "/edit/:quizId": EditQuizView,
    "/create": CreateQuiz,
  };
</script>

<Router {routes} />
