<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
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
        <!-- ---------------------------------------------------------------------------- -->
        <el-row>
          <el-col :span="10">
            <el-form-item label="范围">
              <div class="block">
                <el-date-picker
                        v-model="rangeTimeArray"
                        type="daterange"
                        align="right"
                        unlink-panels
                        range-separator="至"
                        start-placeholder="开始日期"
                        end-placeholder="结束日期"
                        :picker-options="pickerOptions">
                </el-date-picker>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="天数">
              <el-input-number v-model="searchInfo.dayMin" @change="handleChange" :min="1" :max="365" label="最小"></el-input-number>
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
    getAnalysisList
  } from '@/api/stockAnalysis'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Top',
    mixins: [infoList],
    data() {
      return {
        listApi: getAnalysisList,
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
          volume_ratio_max: undefined,
          startTime:undefined,
          endTime:undefined,
          dayMin:1
        },
        rangeTimeArray:undefined
      }
    },
    methods: {
      //搜索
      onSubmit() {
        if (this.rangeTimeArray !== undefined) {
          this.searchInfo.startTime = new Date(this.rangeTimeArray[0]).getTime() / 1000
          this.searchInfo.endTime = new Date(this.rangeTimeArray[1]).getTime() / 1000
        }
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
      //this.getTableData()
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