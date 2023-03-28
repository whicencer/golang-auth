<script lang="ts">
  import { push } from "svelte-spa-router";
  import { toasts, ToastContainer, FlatToast }  from "svelte-toasts";
  import { checkUserNameTaken } from "../lib/checkUsernameTaken";
  import { signup } from "../services/auth/signup";
  
  let username = "";
  let description = "No information given.";
  let password = "";
  
  let userExists = checkUserNameTaken(username) || false;

  const signupHandler = () => {
    signup({
      description,
      password,
      nickname: username
    }).then((response) => {
      if (!response.success) {
        toasts.error(response.message);
      } else {
        toasts.success(response.message);
        push('/#/signin');
      }
    });
  }
</script>

<main>
  <h2>Sign Up</h2>

  <div class="form">
    <input on:input={event => {
      username = event.currentTarget.value;
      checkUserNameTaken(username).then(isTaken => userExists = isTaken).catch(err => err);
    }} type="text" placeholder="Username">
    {#if username.length < 2}
      <span class="error">Username length cannot be less than 2</span>
    {/if}
    {#if userExists === true}
      <span class="error">Username is already taken</span>
    {/if}

    <input on:input={event => {
      description = event.currentTarget.value;
    }} type="text" placeholder="Profile description">
    {#if description.length > 120}
      <span class="error">Description length cannot be more than 120 symbols</span>
    {/if}

    <input on:input={event => {
      password = event.currentTarget.value;
    }} type="password" placeholder="Password">
    {#if password.length < 8}
      <span class="error">Password length cannot be less than 8 symbols</span>
    {/if}
    <button on:click={signupHandler}>Sign up</button>
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