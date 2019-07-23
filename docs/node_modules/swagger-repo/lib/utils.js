const path = require('path');
const fs = require('fs');

const _ = require('lodash');
const glob = require('glob').sync;

exports.pathToFilename = function(path) {
  return path
    .replace(/~1/g, '/')
    .replace(/~0/g, '~')
    .substring(1)
    .replace(/\//g, '@');
};

exports.anyYaml = '**/*.yaml';

exports.getAllYamls = function(dir) {
  return _.fromPairs(
    _.map(glob(path.join(dir, exports.anyYaml)), fname => [path.basename(fname), fname])
  );
};

exports.dirExist = function(path) {
  try {
    return fs.statSync(path).isDirectory();
  } catch (err) {
    if (err && err.code === 'ENOENT') {
      return false;
    }
    throw err;
  }
};
