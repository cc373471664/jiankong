<template>
    <div>
        <a-layout>
            <a-layout-header>
                <a-row>
                    <a-col :span="5">
                        <a-button type="primary" @click="add_click">添加计划</a-button>
                    </a-col>
                </a-row>
            </a-layout-header>
            <a-layout-content>
                <a-table
                        :columns="columns"
                        :dataSource="data"
                        :rowKey="record => record.name"
                        :loading="loading"

                >
                    <span slot="sta" slot-scope="text, t">
                        <template v-if="t.sta==200">   <a-tag color="green">200 👌</a-tag></template>
                          <template v-else><a-tag color="pink">{{t.sta}} 😩</a-tag></template>
                    </span>
                    <span slot="yong" slot-scope="text, t">
                        <template v-if="t.yong"><a-icon @click="click_tingyun(t.name)" type="pause-circle"
                                                        style="color: red;cursor: pointer"/>停止</template>
                        <template v-else><a-icon type="play-circle" @click="click_paoyun(t.name)"
                                                 style="color: #55a532;cursor: pointer"/>启动</template>
                    </span>
                    <span slot="action" slot-scope="text, record">
                        <a-row><a-col :span="8"><a-icon type="stock" @click="rizhi_click(record)"/></a-col>
                            <a-col :span="8"><a-icon @click="del_click(record)" type="delete"/></a-col>
                        </a-row>
        </span>
                </a-table>
            </a-layout-content>
        </a-layout>
        <a-drawer
                title="添加计划"
                :width="720"
                @close="onClose"
                :visible="visible"
                :wrapStyle="{height: 'calc(100% - 108px)',overflow: 'auto',paddingBottom: '108px'}"
        >
            <a-form :form="form">
                <a-form-item label="名字">
                    <a-input v-model="form.name"
                    />
                </a-form-item>
                <a-form-item label="url">
                    <a-input v-model="form.url"/>
                </a-form-item>
                <a-form-item label="间隔(S)">
                    <a-input type="number" v-model.number="form.jiange"/>
                </a-form-item>
            </a-form>
            <a-form-item>
                <a-button @click="handleSubmit" type="primary" html-type="submit">
                    提交
                </a-button>
            </a-form-item>
        </a-drawer>
        <a-drawer
                title="查看日志"
                :width="720"
                @close="onClose2"
                :visible="visible2"
        >
            <a-table
                    :columns="columns2"
                    :dataSource="data2"
                    :pagination="pagination"
                    @change="handleTableChange"
                    :rowKey="record => record.add_time+record.name"
                    :loading="loading"
            >
                 <span slot="sta" slot-scope="text, t">
                         <template v-if="t.sta==200">   <a-tag color="green">200 👌</a-tag></template>
                          <template v-else><a-tag color="pink">{{t.sta}} 😩</a-tag></template>
                    </span>
            </a-table>
        </a-drawer>
    </div>

</template>
<script>
    const columns = [
        {
            title: '名字',
            dataIndex: 'name',
        },
        {
            title: 'url',
            dataIndex: 'url',
        },
        {
            title: '间隔(S)',
            dataIndex: 'jiange',
        },
        {
            title: '最后一次执行时间',
            dataIndex: 'zhi_time',
        },
        {
            title: '最后一次',
            dataIndex: 'sta',
            scopedSlots: {customRender: 'sta'},
        },
        {
            title: '状态',
            dataIndex: 'yong',
            scopedSlots: {customRender: 'yong'},
        },
        {
            title: '操作',
            key: 'caozuo',
            scopedSlots: {customRender: 'action'},
        },

    ];
    const columns2 = [
        {
            title: '时间',
            dataIndex: 'add_time',
        },
        {
            title: '状态',
            dataIndex: 'sta',
            scopedSlots: {customRender: 'sta'},

        },
    ]
    export default {
        mounted() {
            this.fetch();
        },
        created() {
            this.sendWebsocket()
        },
        data() {
            return {
                visible2: false,
                visible: false,
                data: [],
                data2: [],
                loading: false,
                columns,
                columns2,
                form: {"url": "http://"},
                active_name: "",
                pagination: {
                    pageNo: 1,
                    pageSize: 10, // 默认每页显示数量
                    showSizeChanger: true, // 显示可改变每页数量
                    pageSizeOptions: ['10', '20', '50', '100'], // 每页数量选项
                    showTotal: total => `总条数 ${total} `, // 显示总数
                    onShowSizeChange: this.handleTableChange,
                    total: 0 //总条数
                }
            };
        },
        methods: {
            handleTableChange(pagination, filters, sorter) {
                console.log(pagination);
                const pager = {...this.pagination};
                pager.current = pagination.current;
                this.pagination = pager;
                this.request_log(this.active_name, pagination.current, pagination.pageSize)
            },
            fetch() {
                // this.$post("/go/urllist").then((res) => {
                //     this.data = res.data
                // })
            },
            xiu_click(val) {
                console.log(val)
            },
            del_click(val) {
                var that = this
                this.$confirm({
                    title: '确认删除?',
                    content: val.name,
                    okText: 'Yes',
                    okType: 'danger',
                    cancelText: 'No',
                    onOk() {
                        that.$post("/go/dellist", {"name": val.name}).then((res) => {
                            if (res.sta == 1) {
                                that.$message.success(res.data);
                            }
                            that.fetch()
                        })
                    },
                    onCancel() {
                    },
                });
            },
            add_click() {
                this.visible = true
            },
            onClose() {
                this.visible = false

            },
            handleSubmit() {
                if (this.form.url != "" && this.form.name != "" && this.form.jiange != "") {
                    this.$post("/go/addlist", {
                        "url": this.form.url,
                        "name": this.form.name,
                        "jiange": this.form.jiange
                    }).then((res) => {
                        if (res.sta == 1) {
                            this.$message.success(res.data);
                            this.fetch()
                            this.visible = false
                            this.form = {"url": "http://"}
                        } else {
                            this.$message.error(res.data);
                        }

                    })
                } else {
                    this.$message.error("请填写完整")
                }

            },
            onClose2() {
                this.visible2 = false

            },
            request_log(val, current, pagesize) {
                this.data2 = []
                this.$post("/go/urllistlog", {"name": val, "current": current, "pagesize": pagesize}).then((res) => {
                    this.data2 = res.data.data
                    this.pagination.total = res.data.count
                })
            },
            rizhi_click(val) {
                this.active_name = val.name
                this.visible2 = true
                this.request_log(val.name, this.pagination.pageNo, this.pagination.pageSize)
            },
            click_tingyun(name) {
                this.run_req(name, 0)
            },
            click_paoyun(name) {
                this.run_req(name, 1)
            },
            run_req(val, yong) {
                this.$post("/go/send_pao", {"name": val, "yong": yong}).then((res) => {
                    if (res.sta == 0) {
                        this.$message.success(res.data);
                    } else {
                        this.$message.error(res.data)
                    }
                    this.fetch()
                })
            },
            /** 实时更新list数据列表 *******/
            sendWebsocket() {
                var that=this
                let path = window.location.host;
                // let ws = new WebSocket("ws://localhost:8080/go_socket/socketlist");
                let ws=new WebSocket('ws://' + location.host + '/socket/socketlist' )
                //连接打开时触发
                ws.onopen = function (evt) {
                    console.log("打开通道");
                    ws.send("Hello WebSockets!");
                };
                //接收到消息时触发
                ws.onmessage = function (evt) {

                    that.data=JSON.parse(evt.data)
                };
                //连接关闭时触发
                ws.onclose = function (evt) {
                    console.log("Connection closed.");
                };
            }
        },

    };
</script>
