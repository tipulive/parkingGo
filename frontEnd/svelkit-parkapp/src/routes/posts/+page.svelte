<script>
    import {selectedPost} from '$lib/store/store.js'
    // The `data` prop is automatically populated by the `load` function.
    export let data;//to get data from page.server.js
    // @ts-ignore
    function handleClick(PostData){
        selectedPost.set(PostData);
        
        //$inspect(selectedPost);
        //console.log(selectedPost);
    
    }
</script>

<h1>Blog Posts</h1>
<ul>
    {#each data.post as postItem}
        <li on:click={()=>handleClick(postItem)}>{postItem.title}
            
        </li>
    {/each}
</ul>