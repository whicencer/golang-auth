<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { getMe } from "../services/auth/getMe";

  let intervalId;

  let me = null;
  let token = localStorage.getItem('authToken');
  let isAuthorized = !!token.length;

  const handleStorageChange = () => {
    token = localStorage.getItem('authToken');
    if (!token.length) {
      isAuthorized = false;
    } else {
      isAuthorized = true;
    }
  }

  if (isAuthorized) {
    getMe(token)
      .then(data => {
        if (data.success) {
          me = { username: data.nickname, description: data.description };
        } else {
          console.log('Not authenticated');
        }
      })
  }

  onMount(() => {
    intervalId = setInterval(() => {
      handleStorageChange();
    }, 500);
  });

  onDestroy(() => {
    clearInterval(intervalId);
  });
</script>

<header>
  <h2>
    <a href="#/">Golang auth</a>
  </h2>
  {#if isAuthorized && me !== null}
    <h3>{me.username}</h3>
  {:else if !isAuthorized}
    <nav>
      <a href="#/signin">Sign In</a>
      <a href="#/signup">Sign Up</a>
    </nav>
  {/if}
</header>

<style>
  header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  nav > a {
    margin: 0 10px;
  }

  h2 > a {
    color: inherit;
  }
</style>