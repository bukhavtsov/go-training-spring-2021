## Task 5

## Description

You need to develop the image that allows to convert HTML to (Github) Markdown.

1. Start a container with bash in an interactive mode using `ubuntu:latest` as the base image
2. Install `curl` inside the container
3. Install the latest version of pandoc from official releases:
   `https://github.com/jgm/pandoc/releases`
4. Try downloading any page, say `https://google.com` using `curl`, and piping it through
   pandoc. Use arguments such that pandoc will read HTML and convert it to Github
   Flavoured Markdown
5. Write a Dockerfile which will create an image from `ubuntu:latest`, that will contain `curl`
   and pandoc, and will run the `curl | pandoc | less` pipeline. The container run from that
   image must take at least one argument, URI, which must be passed as an argument to
   `curl` in the pipeline. This way you should get an image of a primitive text-based internet
   browser
6. Build an image from the `Dockerfile`
7. Authenticate with your Docker Hub account’s credentials in your Docker client
8. Push the built image into the Docker Hub so anyone could pull it
9. Create `README.md` file with instructions on how to get and run the containerized
    application (via Dockerfile and Docker hub)
10. Create a private Github repository
11. Add me to the list of contributors
12. Add the Dockerfile and other supplementary files you’ve created
13. Make commit
14. Push the branch into the Github repository
15. Create PR. Add me to the list of reviewers. Wait for the PR to get reviewed and merge it.

#### Let me know if you have any questions ;)
