

var exec = require('child_process').execSync; 
var exeFile = '/Users/luffyyzhang/WorkSpace/tabtoy/tabtoy'
var globalExcel = './excel/GlobalDefine/Globals.xlsx'
var excelToJsonCmd = exeFile + ' --mode=exportorv2 --json_out=%s %s %s'

var util = require('util');
var fs= require("fs");
var rd = require("rd");
var path = require("path");

rd.eachSync("./excel", function (f, s) {
    // 每找到一个文件都会调用一次此函数
    // 参数s是通过 fs.stat() 获取到的文件属性值
    if (f.endsWith(".xlsx")) {
        if (f.endsWith("Globals.xlsx")) {
            return;
        }
        console.log('file: %s', f);
        var myPath = f.replace("excel", "json")
        var pathobj = path.parse(myPath);

        if (!fs.existsSync(pathobj.dir)) {
            fs.mkdirSync(pathobj.dir);
        }

        var destFile = pathobj.dir + "/" + pathobj.name + ".json";

        var cmdStr = util.format(excelToJsonCmd, destFile, globalExcel, f);
        //console.log("cmd:", cmdStr);
        exec(cmdStr, { stdio: [null, process.stdout, process.stderr] });
    }
});

var cmdStr = 'node json-to-go.js ./json entity';
exec(cmdStr, { stdio: [null, process.stdout, process.stderr] });

console.log("All 操作结束！");
