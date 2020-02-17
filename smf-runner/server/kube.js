const { exec } = require("child_process");
function getClusters() {
  return new Promise((resolve, reject) => {
    exec("kubectl config get-contexts -o=name", function(err, data) {
      if (err) {
        reject(err);
        return;
      }
      resolve(
        data
          .toString()
          .split("\n")
          .filter(Boolean)
      );
    });
  });
}

module.exports = { getClusters };
