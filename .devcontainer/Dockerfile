FROM mcr.microsoft.com/devcontainers/base:jammy

# [Option] Install zsh
ARG INSTALL_ZSH="false"
# [Option] Upgrade OS packages to their latest versions
ARG UPGRADE_PACKAGES="false"
# Install needed packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

COPY library-scripts/*.sh /tmp/library-scripts/

RUN apt-get update 
RUN  /bin/bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" "false" "true" 
RUN  apt-get install -y curl 
RUN  apt-get clean -y 
RUN  rm -rf /var/lib/apt/lists/* /tmp/library-scripts

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
# install swagger
ADD getswagger.sh /usr/local/bin
RUN /usr/local/bin/getswagger.sh