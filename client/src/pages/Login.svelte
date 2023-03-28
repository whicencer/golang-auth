<script lang="ts">
  import { signin } from "../services/auth/signin";
  import { toasts, ToastContainer, FlatToast }  from "svelte-toasts";
  import { push } from "svelte-spa-router";

  let username = '';
  let password = '';

  const signInHandler = () => {
    signin({
      nickname: username,
      password
    }).then((data) => {
      if (data.success) {
        toasts.success(data.message);
        localStorage.setItem('authToken', data.token);
        push('/');
      } else {
        toasts.error(data.message);
      }
    })
  }
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