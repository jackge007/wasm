FROM l.gcr.io/google/bazel:1.2.1

RUN apt update && \
    apt install bzip2 libxml2 -y

RUN curl -sL https://deb.nodesource.com/setup_13.x | bash - && \
    apt install -y nodejs

RUN apt install expect-dev -y

RUN npm install -g @bazel/bazelisk

# TODO: use the latest stable release package after the WASI supported is officially released
#   this corresponds to https://github.com/tinygo-org/tinygo/commit/9a015f4f6441fd68eabe35b3dc9748a5467d5934
RUN wget https://storage.googleapis.com/getenvoy-package/tinygo-nightly-packages/tinygo_amd64.deb
RUN dpkg -i tinygo_amd64.deb && rm tinygo_amd64.deb

RUN wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz && rm go1.15.2.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# libstdc++6.so is required for TinyGo
RUN add-apt-repository ppa:ubuntu-toolchain-r/test && \
    apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y dist-upgrade && \
    rm -rf /var/lib/apt/lists/*

COPY build-filter.sh /build-filter.sh

ENTRYPOINT /build-filter.sh
