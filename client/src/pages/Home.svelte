<script>
  const getMe = async () => {
    const response = await fetch('http://127.0.0.1:2000/auth/me', {
      headers: new Headers({
        'Authorization': `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXNjcmlwdGlvbiI6IkhlbGxvLCBteSBuYW1lIGlzIERlbmllbCwgSSdtIGEgcHJvZ3JhbW1lciIsImV4cCI6MTY4MDE3NDI0Niwibmlja25hbWUiOiJ3aGljZW5jZXIifQ._oZkn5KkIq4yt3bcIY0hEQ9uu0mXsu09qTdZXagR5zM`
      })
    });
    const data = await response.json();

    console.log(data);
  }
</script>

<main>
  <h2>Home</h2>
  <button on:click={getMe}>Get me</button>
</main>