async function increment() {
  await ddClient.extension.vm.cli.exec("/counter", [], {
    stream: {
      onOutput(data) {
        if (data.stdout) {
          ddClient.desktopUI.toast.success(data.stdout);
        } else {
          ddClient.desktopUI.toast.error(data.stderr);
        }
      },
      onError(error) {
        ddClient.desktopUI.toast.error(error);
      },
      onClose(exitCode) {
        console.log("onClose with exit code " + exitCode);
      },
    },
  });
}
