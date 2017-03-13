#!/usr/bin/env node
/*
 * Usage: ./filljson.js valueType <template.json "path.to.target" <(some command)
 * valueType: [string] string int float
 */

var fs = require('fs');

var valueType = process.argv[2]
var paths = process.argv[3].split('.');
var dataFile = process.argv[4];

function parseStringArray(fileBody) {
    return fileBody
      .split('\n')
    .filter(function(line) {
        return line !== '';
    });
}

function parseString(fileBody) {
    return fileBody;
}

var parsers = {
    '[string]': parseStringArray,
    'int': parseInt,
    'float': parseFloat,
    'string': parseString,
}
var parser = parsers[valueType] || parseString;

var json = JSON.parse(require('fs').readFileSync('/dev/stdin', 'utf8'));
var value = parser(fs.readFileSync(dataFile, 'utf8'));

var finger = json;
for (var idx in paths) {
    var key = paths[idx];
    if (idx == paths.length - 1) {
        finger[key] = value;
    } else {
        finger = finger[key];
    }
}
console.log(JSON.stringify(json));
