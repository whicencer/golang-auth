<script lang="ts">
  import { push } from "svelte-spa-router";
  import { toasts, FlatToast, ToastContainer } from "svelte-toasts";
  import { onMount } from 'svelte';

  let isAuthorized = false;
	onMount(async () => {
    isAuthorized = localStorage.getItem('authToken').length !== 0;
	});


  const logout = () => {
    localStorage.setItem('authToken', '');
    push('/signin')
    toasts.info('You have successfully logged out');
  }
</script>

<main>
  <h2>Home</h2>
  <div>
    {#if isAuthorized}
      <p>You logged in</p>
      <button on:click={logout}>Logout</button>
    {:else}
      <p>You're not logged in</p>  
    {/if}
  </div>
  <ToastContainer let:data={data}>
    <FlatToast {data} />
  </ToastContainer>
</main>

<style>
  main > div {
    max-width: 320px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
  }
</style>