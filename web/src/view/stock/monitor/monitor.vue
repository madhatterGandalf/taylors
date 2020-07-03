<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-row>
          <el-col :span="10">
            <el-form-item label="名称">
              <el-input placeholder="平安银行" v-model="searchInfo.name" clearable ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="编码">
              <el-input placeholder="SZ0000001"  v-model="searchInfo.code" clearable ></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="市值">
              <el-input-number placeholder="最小" v-model="searchInfo.marketCapitalMin" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.marketCapitalMax" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="股价">
              <el-input-number placeholder="最小" v-model="searchInfo.currentMin" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.currentMax" ></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="涨幅">
              <el-input-number placeholder="最小"  v-model="searchInfo.percentMin" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.percentMax" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="量比">
              <el-input-number placeholder="最小"  v-model="searchInfo.volume_ratio_min" ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.volume_ratio_max" ></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('add')" type="primary">新增监控</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="downBek" type="primary">下载</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe :default-sort = "{prop: ['f20','f2','f3','f10','f5','f6'], order: 'descending'}">
      <el-table-column label="名称" min-width="70" prop="f14"></el-table-column>
      <el-table-column label="编码" min-width="70" prop="f12"></el-table-column>
      <el-table-column label="市值" min-width="80" prop="f20" sortable ></el-table-column>
      <el-table-column label="当前价" min-width="80" prop="f2" sortable ></el-table-column>

      <el-table-column label="涨幅" min-width="80" prop="f3" sortable></el-table-column>
      <el-table-column label="量比" min-width="80" prop="f10" sortable></el-table-column>

      <el-table-column label="成交量" min-width="120" prop="f5" sortable></el-table-column>
      <el-table-column label="成交额" min-width="120" prop="f6" sortable></el-table-column>
      <el-table-column label="高位预警" min-width="120" prop="monitor_high" sortable></el-table-column>
      <el-table-column label="低位预警" min-width="120" prop="monitor_low" sortable></el-table-column>
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editStockMonitor(scope.row)" size="small"  style="float: right;" type="primary" >编辑</el-button>
          <el-button @click="deleteStockMonitor(scope.row)" size="small" style="float: left;" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="stockMonitorForm">
        <el-form-item label="编码" prop="code">
          <el-input autocomplete="off" v-model="form.code" clearable ></el-input>
        </el-form-item>
        <el-form-item label="高位预警" prop="monitor_high">
          <el-input-number placeholder="1.2" v-model="form.monitor_high" :controls="false"></el-input-number>
        </el-form-item>
        <el-form-item label="低位预警" prop="monitor_low">
          <el-input-number placeholder="1.2" v-model="form.monitor_low" :controls="false"></el-input-number>
        </el-form-item>
      </el-form>
<!--      <div class="warning">新增监控需要在角色管理内配置权限才可使用</div>-->
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>


<script>
  import {
    getMonitorOne,
    getMonitorList,
    updateMonitor,
    delMonitor,
    addMonitor,
  } from '@/api/stockMonitor'
  import infoList from '@/components/mixins/infoList'
  const methodOptions = [
    {
      value: 'POST',
      label: '创建',
      type: 'success'
    },
    {
      value: 'GET',
      label: '查看',
      type: ''
    },
    {
      value: 'PUT',
      label: '更新',
      type: 'warning'
    },
    {
      value: 'DELETE',
      label: '删除',
      type: 'danger'
    }
  ]
  export default {
    name: 'Top',
    mixins: [infoList],
    data() {
      return {
        dialogTitle: '新增监控',
        form: {
          code: '',
          monitor_high: 0,
          monitor_low: 0
        },
        dialogFormVisible: false,
        listApi: getMonitorList,
        searchInfo: {
          name:undefined,
          code:undefined,
          currentMax:undefined,
          currentMin:undefined,
          marketCapitalMin: undefined,
          marketCapitalMax: undefined,
          percentMin: undefined,
          percentMax: undefined,
          volume_ratio_min: undefined,
          volume_ratio_max: undefined
        },
        methodOptions: methodOptions,
        type: '',
      }
    },
    methods: {
      //搜索
      onSubmit() {
        this.getTableData()
      },

      downBek() {
        var bodyStr = "\r\n"
        var data = this.tableData
        for (var i =0; i<data.length; i++){
          bodyStr += data[i].f12 + "\r\n"
        }

        const element = document.createElement('a')
        element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(bodyStr))
        element.setAttribute('download', "top.BEK")
        element.style.display = 'none'
        element.click()
      },

      openDialog(type) {
        switch (type) {
          case 'add':
            this.dialogTitle = '新增监控'
            break
          case 'edit':
            this.dialogTitle = '编辑监控'
            break
          default:
            break
        }
        this.type = type
        this.dialogFormVisible = true
      },
      async enterDialog() {
        this.$refs.stockMonitorForm.validate(async valid => {
          if (valid) {
            if (this.form.monitor_high <= this.form.monitor_low){
              this.$message({
                type: 'error',
                message: '参数错误',
                showClose: true
              })
              return
            }
            switch (this.type) {
              case 'add':
              {
                const res = await addMonitor(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }

                break
              case 'edit':
              {
                const res = await updateMonitor(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
                break
              default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
                break
            }
          }
        })
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      async editStockMonitor(row) {
        const res = await getMonitorOne({ id: row.id })
        this.form = res.data
        this.openDialog('edit')
      },
      initForm() {
        this.$refs.stockMonitorForm.resetFields()
        this.form= {
          code: '',
          monitor_high: undefined,
          monitor_low: undefined
        }
      },
      async deleteStockMonitor(row) {
        this.$confirm('此操作将永久删除当前监控股票, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
            const res = await delMonitor(row)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '删除成功!'
              })
              this.getTableData()
            }
          })
          .catch(() => {
            this.$message({
              type: 'info',
              message: '已取消删除'
            })
          })
      }
    },
    created(){
      this.getTableData()
      // setInterval(()=>{
      //   this.getTableData()
      // },20000)
    }
  }
</script>
<style scoped lang="scss">
  .button-box {
    padding: 10px 20px;
  }
  .el-button {
    float: right;
  }

  .el-tag--mini {
    margin-left: 5px;
  }
  .warning {
    color: #dc143c;
  }
</style>