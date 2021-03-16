# igo

## 环境说明

1. 去掉操作系统的GOPATH 环境变量
2. 安装The VS Code Go extension
3. 修改配置文件增加如下配置

    ```
        "go.toolsGopath": "/TOOLS/GOPATH",
        "go.goroot": "/TOOLS/software/go",
        "gopls": {
                "experimentalWorkspaceModule": true,
        }
    ```

4. 需要先安装号gotools
