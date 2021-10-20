PACKER_VERSION=1.7.6
BADGE_HOSTNAME=badge
BADGE_VERSION=master

all: clean install image

install:
	sudo apt update -y
	sudo apt install qemu-system-arm qemu-user-static binfmt-support bmap-tools kpartx unzip zip tar -y -qq
	curl -O "https://golang.org/dl/go1.17.2.linux-amd64.tar.gz"
	tar -C /usr/local/ -xf go1.17.2.linux-amd64.tar.gz
	ln -s /usr/local/go/bin/go /usr/bin/go
	rm go1.17.2.linux-amd64.tar.gz
	curl --silent https://releases.hashicorp.com/packer/$(PACKER_VERSION)/packer_$(PACKER_VERSION)_linux_amd64.zip -o /tmp/packer.zip
	unzip /tmp/packer.zip -d /tmp
	sudo mv /tmp/packer /usr/bin/packer
	git clone https://github.com/solo-io/packer-builder-arm-image /tmp/packer-builder-arm-image
	cd /tmp/packer-builder-arm-image && go get -d ./... && go build
	sudo cp /tmp/packer-builder-arm-image/packer-builder-arm-image /usr/bin

image:
	cd builder && sudo /usr/bin/packer build -var "badge_hostname=$(BADGE_HOSTNAME)" -var "badge_version=$(BADGE_VERSION)" badge.json
	sudo mv builder/output-badge/image badge-raspbian-lite-$(BADGE_VERSION).img
	sudo sha256sum badge-raspbian-lite-$(BADGE_VERSION).img > badge-raspbian-lite-$(BADGE_VERSION).sha256
	# sudo zip badge-raspbian-lite-$(BADGE_VERSION).zip badge-raspbian-lite-$(BADGE_VERSION).sha256 badge-raspbian-lite-$(BADGE_VERSION).img

clean:
	rm -rf /tmp/packer-builder-arm-image
	rm -f badge-raspbian-lite-*.zip badge-raspbian-lite-*.img badge-raspbian-lite-*.sha256
	rm -rf builder/output-badge  builder/packer_cache
