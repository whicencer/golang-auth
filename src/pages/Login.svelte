<script lang="ts">
  import { signin } from "../services/auth/signin";
  import { toasts, ToastContainer, FlatToast }  from "svelte-toasts";
  import { push } from "svelte-spa-router";
  import { isAuthorized } from "../store";

  let username = '';
  let password = '';

  const signInHandler = async () => {
    const response = await signin({
      nickname: username,
      password
    });

    try {
      if (!response.ok) {
        throw new Error(response.message);
      }
      toasts.success(response.message);
      localStorage.setItem('authToken', response.token);
      push('/');
    } catch (error) {
      toasts.error(error.message);
    }
  }

  isAuthorized.subscribe(value => {
    if (value.isAuthorized)
      push('/');
  })
</script>

<main>
  <h2>Login</h2>
  <div class="form">
    <input on:input={event => {
      username = event.currentTarget.value
    }}
    type="text"
    placeholder="Username">
  
    <input on:input={event => {
      password = event.currentTarget.value
    }}
    type="password"
    placeholder="Password">
  
    <button on:click={signInHandler}>Sign in</button>
  </div>
  <ToastContainer let:data={data}>
		<FlatToast {data}  />
	</ToastContainer>
</main>

<style>
  button {
    margin-top: 10px;
  }
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
</style>