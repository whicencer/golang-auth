<script>
  import Router from 'svelte-spa-router'
  import Header from '../components/Header.svelte';
  import Home from '../pages/Home.svelte'
  import Login from '../pages/Login.svelte';
  import Register from '../pages/Register.svelte';
</script>

<main>
  <Header />
  <Router routes={{
    '/': Home,
    '/signup': Register,
    '/signin': Login,
  }} />
</main>