<script lang="ts">
  import { onMount } from "svelte";
  import { push } from "svelte-spa-router";
  import { toasts, FlatToast, ToastContainer } from "svelte-toasts";
  import { getMe } from "../services/auth/getMe";
  import { isAuthorized } from "../store";

  onMount(async () => {
    const token = localStorage.getItem('authToken');
    if (token.length !== 0) {
      const response = await getMe(token);

      try {
        if (!response.ok) {
          throw new Error(response.message);
        }
        isAuthorized.set({
          isAuthorized: true,
          me: { description: response.description, username: response.nickname }
        });
      } catch (error) {
        toasts.error(error.message);
      }
    }
  });

  let isAuthorizedValue = false;
  
  isAuthorized.subscribe(value => {
    if (!value.isAuthorized)
      push('/signin');
		isAuthorizedValue = value.isAuthorized;
	});

  const logout = () => {
    localStorage.setItem('authToken', '');
    isAuthorized.set({
      isAuthorized: false,
      me: null
    });
    push('/signin')
    toasts.info('You have successfully logged out');
  }
</script>

<main>
  <h2>Home</h2>
  <div>
    {#if isAuthorizedValue}
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