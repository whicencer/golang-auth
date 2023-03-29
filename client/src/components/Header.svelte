<script lang="ts">
  import { isAuthorized } from "../store";

  let me = null;
  let isAuthorizedValue = false;
  
  isAuthorized.subscribe(value => {
		isAuthorizedValue = value.isAuthorized;
    me = value.me;
	});
</script>

<header>
  <h2>
    <a href="#/">Golang auth</a>
  </h2>
  {#if isAuthorizedValue && me !== null}
    <h3>{me.username}</h3>
  {:else if !isAuthorizedValue}
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