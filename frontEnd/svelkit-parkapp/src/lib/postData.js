// This could be fetched from a database or a CMS in a real app.
export const posts = [
    {
        id: 1,
        slug: 'getting-started-with-sveltekit',
        title: 'Getting Started with SvelteKit',
        content: 'SvelteKit is a framework for building web applications of all sizes...'
    },
    {
        id: 2,
        slug: 'understanding-load-functions',
        title: 'Understanding Load Functions',
        content: 'Load functions are the heart of data fetching in SvelteKit...'
    },
    {
        id: 3,
        slug: 'advanced-routing-patterns',
        title: 'Advanced Routing Patterns',
        content: 'SvelteKit\'s file-based router is powerful and flexible...'
    }
];

// A helper function to find a single post by its slug
// @ts-ignore
export function getPostBySlug(slug) {
    return posts.find(post => post.slug === slug);
}