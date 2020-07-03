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
              <el-input-number placeholder="最小" v-model="searchInfo.marketCapitalMin"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.marketCapitalMax"  ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="股价">
              <el-input-number placeholder="最小" v-model="searchInfo.currentMin"  ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.currentMax"  ></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="涨幅">
              <el-input-number placeholder="最小"  v-model="searchInfo.percentMin"  ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.percentMax"  ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="量比">
              <el-input-number placeholder="最小"  v-model="searchInfo.volume_ratio_min"  ></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.volume_ratio_max"  ></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="downBek" type="primary">下载</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe :default-sort = "{prop: ['f20','f2','f3','f10','f5','f6'], order: 'descending'}">
      <el-table-column label="名称" min-width="50" prop="f14"></el-table-column>
      <el-table-column label="编码" min-width="50" prop="f12"></el-table-column>

      <el-table-column label="市值" min-width="70" prop="f20" sortable ></el-table-column>
      <el-table-column label="当前价" min-width="70" prop="f2" sortable></el-table-column>

      <el-table-column label="涨幅" min-width="80" prop="f3" sortable></el-table-column>
      <el-table-column label="量比" min-width="80" prop="f10" sortable></el-table-column>

      <el-table-column label="成交量" min-width="120" prop="f5" sortable></el-table-column>
      <el-table-column label="成交额" min-width="120" prop="f6" sortable></el-table-column>

      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="addStockMonitorDay(scope.row)" size="small" type="primary">监控</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :style="{float:'right',padding:'20px'}"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>


<script>
  import {
    addMonitor
  } from '@/api/stockMonitor'
  import {
    getAllList
  } from '@/api/stockAll'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Top',
    mixins: [infoList],
    data() {
      return {
        listApi: getAllList,
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
      }
    },
    methods: {
      //搜索
      onSubmit() {
        this.page = 1
        this.pageSize = 10
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

      async addStockMonitorDay(row) {
        const res = await addMonitor({isDay:false,code: row.f12})
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '监控成功!'
          })
        }else{
          this.$message({
            type: 'error',
            message: '监控失败!'
          })
        }
      }
    },
    created(){
      this.getTableData()
      // setInterval(()=>{
      //   this.getTableData()
      // },10000)
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