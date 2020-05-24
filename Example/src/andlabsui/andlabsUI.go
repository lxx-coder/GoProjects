package andlabsui

import (
    "github.com/andlabs/ui"
    _ "github.com/andlabs/ui/winmanifest"
)

func andlabsUI() {
    err := ui.Main(func() {
        // 生成：文本框
        //name := ui.NewEntry()
        name := ui.NewMultilineEntry()
        //name := ui.NewNonWrappingMultilineEntry()

        // 生成：标签
        //greeting := ui.NewLabel("欢迎")
        greeting := ui.NewMultilineEntry()
        // 生成按钮
        button := ui.NewButton("欢迎")
        button1 := ui.NewButton("hello")
        // 设置：按钮点击事件
        button.OnClicked(func(*ui.Button) {
            greeting.SetText("你好," + name.Text() + "!")
        })
        button1.OnClicked(func(*ui.Button) {
            name.SetText("hello," + greeting.Text() + "!")
        })
        // 生成垂直容器
        box := ui.NewVerticalBox()

        // 往垂直容器中添加控件
        box.Append(ui.NewLabel("请输入你的名字："), false)
        box.Append(name, true)
        box.Append(button, false)
        box.Append(button1, false)
        box.Append(greeting, true)

        //box1 := ui.NewHorizontalBox()
        //box1.Append(button,false)
        //box1.Append(button1,false)


        // 生成：窗口（标题，宽度，高度，是否有菜单控件）
        window := ui.NewWindow("您好", 800, 400, false)

        // 窗口容器绑定
        window.SetChild(box)
        //window.SetChild(box1)

        //设置：窗口关闭时
        window.OnClosing(func(*ui.Window) bool {
            // 窗口关闭
            ui.Quit()
            return true
        })

        // 窗口显示
        window.Show()
    })

    if err != nil {
        panic(err)
    }
}
