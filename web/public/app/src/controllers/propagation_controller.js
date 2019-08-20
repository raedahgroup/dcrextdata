import { Controller } from 'stimulus'
import axios from 'axios'
import { hide, show, setActiveOptionBtn, legendFormatter, showLoading, hideLoading, displayPillBtnOption, setActiveRecordSetBtn } from '../utils'

const Dygraph = require('../../../dist/js/dygraphs.min.js')

export default class extends Controller {
  static get targets () {
    return [
      'nextPageButton', 'previousPageButton', 'recordSetSelector', 'bothRecordSetOption', 'selectedRecordSet', 'selectedNum', 'numPageWrapper', 'paginationButtonsWrapper',
      'tablesWrapper', 'table', 'blocksTbody', 'votesTbody', 'chartWrapper', 'chartsView', 'labels', 'axisOption',
      'blocksTable', 'blocksTableBody', 'blocksRowTemplate', 'votesTable', 'votesTableBody', 'votesRowTemplate',
      'totalPageCount', 'currentPage', 'viewOptionControl', 'chartSelector', 'viewOption', 'loadingData'
    ]
  }

  initialize () {
    this.currentPage = parseInt(this.currentPageTarget.getAttribute('data-current-page'))
    if (this.currentPage < 1) {
      this.currentPage = 1
    }

    this.selectedRecordSetTargets.forEach(li => {
      if (this.selectedViewOption === 'table' && li.dataset.option === 'both') {
        li.classList.add('active')
      } else {
        li.classList.remove('active')
      }
    })

    this.selectedViewOption = this.viewOptionControlTarget.getAttribute('data-initial-value')
    if (this.selectedViewOption === 'chart') {
      this.setChart()
    } else {
      this.setTable()
    }
  }

  setTable () {
    this.selectedViewOption = 'table'
    this.selectedRecordSet = 'both'
    setActiveOptionBtn(this.selectedViewOption, this.viewOptionTargets)
    hide(this.chartWrapperTarget)
    show(this.paginationButtonsWrapperTarget)
    show(this.numPageWrapperTarget)
    hide(this.chartWrapperTarget)
    show(this.tablesWrapperTarget)
    setActiveRecordSetBtn(this.selectedRecordSet, this.selectedRecordSetTargets)
    displayPillBtnOption(this.selectedViewOption, this.selectedRecordSetTargets)
    this.fetchTableData(this.currentPage)
  }

  setChart () {
    this.selectedViewOption = 'chart'
    this.selectedRecordSet = 'blocks'
    hide(this.numPageWrapperTarget)
    hide(this.paginationButtonsWrapperTarget)
    hide(this.tablesWrapperTarget)
    show(this.chartWrapperTarget)
    setActiveOptionBtn(this.selectedViewOption, this.viewOptionTargets)
    setActiveRecordSetBtn(this.selectedRecordSet, this.selectedRecordSetTargets)
    displayPillBtnOption(this.selectedViewOption, this.selectedRecordSetTargets)
    this.fetchChartDataAndPlot()
  }

  setBothRecordSet () {
    this.selectedRecordSet = 'both'
    setActiveOptionBtn(this.selectedRecordSet, this.selectedRecordSetTargets)
    this.currentPage = 1
    this.selectedNumTarget.value = this.selectedNumTarget.options[0].text
    if (this.selectedViewOption === 'table') {
      this.fetchTableData(1)
    } else {
      this.fetchChartDataAndPlot()
    }
  }

  setBlocksRecordSet () {
    this.selectedRecordSet = 'blocks'
    setActiveOptionBtn(this.selectedRecordSet, this.selectedRecordSetTargets)
    this.currentPage = 1
    this.selectedNumTarget.value = this.selectedNumTarget.options[0].text
    if (this.selectedViewOption === 'table') {
      this.fetchTableData(1)
    } else {
      this.fetchChartDataAndPlot()
    }
  }

  setVotesRecordSet () {
    this.selectedRecordSet = 'votes'
    setActiveOptionBtn(this.selectedRecordSet, this.selectedRecordSetTargets)
    this.currentPage = 1
    this.selectedNumTarget.value = this.selectedNumTarget.options[0].text
    if (this.selectedViewOption === 'table') {
      this.fetchTableData(1)
    } else {
      this.fetchChartDataAndPlot()
    }
  }

  loadPreviousPage () {
    this.fetchTableData(this.currentPage - 1)
  }

  loadNextPage () {
    this.fetchTableData(this.currentPage + 1)
  }

  numberOfRowsChanged () {
    this.selectedNum = this.selectedNumTarget.value
    this.fetchTableData(1)
  }

  setAxis (e) {
    var target = e.srcElement || e.target
    this.selectedAxis = target ? target.dataset.option : e
    setActiveOptionBtn(this.selectedAxis, this.axisOptionTargets)
    this.fetchChartDataAndPlot()
  }

  fetchTableData (page) {
    const _this = this

    let elementsToToggle = [this.tablesWrapperTarget]
    showLoading(this.loadingDataTarget, elementsToToggle)

    var numberOfRows = this.selectedNumTarget.value
    let url = '/getpropagationdata'
    switch (this.selectedRecordSet) {
      case 'blocks':
        url = 'getblocks'
        break
      case 'votes':
        url = 'getvotes'
        break
      default:
        url = 'getpropagationdata'
        break
    }
    axios.get(`/${url}?page=${page}&records-per-page=${numberOfRows}&view-option=${_this.selectedViewOption}`).then(function (response) {
      hideLoading(_this.loadingDataTarget, elementsToToggle)
      let result = response.data
      _this.totalPageCountTarget.textContent = result.totalPages
      _this.currentPageTarget.textContent = result.currentPage
      const pageUrl = `propagation?page=${result.currentPage}&records-per-page=${result.selectedNum}&record-set=${_this.selectedRecordSet}&view-option=${_this.selectedViewOption}`
      window.history.pushState(window.history.state, _this.addr, pageUrl)

      _this.currentPage = result.currentPage
      if (_this.currentPage <= 1) {
        hide(_this.previousPageButtonTarget)
      } else {
        show(_this.previousPageButtonTarget)
      }

      if (_this.currentPage >= result.totalPages) {
        hide(_this.nextPageButtonTarget)
      } else {
        show(_this.nextPageButtonTarget)
      }

      _this.displayData(result)
    }).catch(function (e) {
      hideLoading(_this.loadingDataTarget, elementsToToggle)
      console.log(e) // todo: handle error
    })
  }

  displayData (data) {
    switch (this.selectedRecordSet) {
      case 'blocks':
        this.displayBlocks(data.records)
        break
      case 'votes':
        this.displayVotes(data.voteRecords)
        break
      default:
        this.displayPropagationData(data.records)
        break
    }
  }

  displayBlocks (data) {
    const _this = this
    this.blocksTableBodyTarget.innerHTML = ''
    data.forEach(block => {
      const exRow = document.importNode(_this.blocksRowTemplateTarget.content, true)
      const fields = exRow.querySelectorAll('td')

      fields[0].innerHTML = `<a target="_blank" href="https://explorer.dcrdata.org/block/${block.block_height}">${block.block_height}</a>`
      fields[1].innerText = block.block_internal_time
      fields[2].innerText = block.block_receive_time
      fields[3].innerText = block.delay
      fields[4].innerHTML = `<a target="_blank" href="https://explorer.dcrdata.org/block/${block.block_height}">${block.block_hash}</a>`

      _this.blocksTableBodyTarget.appendChild(exRow)
    })

    hide(this.tableTarget)
    hide(this.votesTableTarget)
    show(this.blocksTableTarget)
  }

  displayVotes (data) {
    const _this = this
    this.votesTableBodyTarget.innerHTML = ''

    if (data) {
      data.forEach(item => {
        const exRow = document.importNode(_this.votesRowTemplateTarget.content, true)
        const fields = exRow.querySelectorAll('td')

        fields[0].innerHTML = `<a target="_blank" href="https://explorer.dcrdata.org/block/${item.voting_on}">${item.voting_on}</a>`
        fields[1].innerHTML = `<a target="_blank" href="https://explorer.dcrdata.org/block/${item.block_hash}">...${item.short_block_hash}</a>`
        fields[2].innerText = item.validator_id
        fields[3].innerText = item.validity
        fields[4].innerText = item.receive_time
        fields[5].innerText = item.block_time_diff
        fields[6].innerText = item.block_receive_time_diff
        fields[7].innerHTML = `<a target="_blank" href="https://explorer.dcrdata.org/tx/${item.hash}">${item.hash}</a>`

        _this.votesTableBodyTarget.appendChild(exRow)
      })
    } else {
      this.votesTableBodyTarget.innerHTML = 'No votes to show'
    }

    hide(this.tableTarget)
    hide(this.blocksTableTarget)
    show(this.votesTableTarget)
  }

  displayPropagationData (data) {
    let blocksHtml = ''
    data.forEach(block => {
      let votesHtml = ''
      let i = 0
      if (block.votes) {
        block.votes.forEach(vote => {
          votesHtml += `<tr>
                            <td><a target="_blank" href="https://explorer.dcrdata.org/block/${vote.voting_on}">${vote.voting_on}</a></td>
                            <td><a target="_blank" href="https://explorer.dcrdata.org/block/${vote.block_hash}">...${vote.short_block_hash}</a></td>
                            <td>${vote.validator_id}</td>
                            <td>${vote.validity}</td>
                            <td>${vote.receive_time}</td>
                            <td>${vote.block_receive_time_diff}s</td>
                            <td><a target="_blank" href="https://explorer.dcrdata.org/tx/${vote.hash}">${vote.hash}</a></td>
                        </tr>`
        })
      }

      let padding = i > 0 ? 'style="padding-top:50px"' : ''
      i++
      blocksHtml += `<tbody data-target="propagation.blockTbody"
                            data-block-hash="${block.block_hash}">
                        <tr>
                            <td colspan="100" ${padding}>
                              <span class="d-inline-block"><b>Height</b>: ${block.block_height} </span>  &#8195;
                              <span class="d-inline-block"><b>Timestamp</b>: ${block.block_internal_time}</span>  &#8195;
                              <span class="d-inline-block"><b>Received</b>: ${block.block_receive_time}</span>  &#8195;
                              <span class="d-inline-block"><b>Hash</b>: <a target="_blank" href="https://explorer.dcrdata.org/block/${block.block_height}">${block.block_hash}</a></span>
                            </td>
                        </tr>
                        </tbody>
                        <tbody data-target="propagation.votesTbody" data-block-hash="${block.block_hash}">
                        <tr style="white-space: nowrap;">
                            <td style="width: 120px;">Voting On</td>
                            <td style="width: 120px;">Block Hash</td>
                            <td style="width: 120px;">Validator ID</td>
                            <td style="width: 120px;">Validity</td>
                            <td style="width: 120px;">Received</td>
                            <td style="width: 120px;">Block Receive Time Diff</td>
                            <td style="width: 120px;">Hash</td>
                        </tr>
                        ${votesHtml}
                        </tbody>
                          <tr>
                              <td colspan="7" height="15" style="border: none !important;"></td>
                          </tr>`
    })

    this.tableTarget.innerHTML = blocksHtml

    show(this.tableTarget)
    hide(this.blocksTableTarget)
    hide(this.votesTableTarget)
  }

  fetchChartDataAndPlot () {
    let elementsToToggle = [this.chartWrapperTarget]
    showLoading(this.loadingDataTarget, elementsToToggle)

    const _this = this
    axios.get(`/propagationchartdata?record-set=${this.selectedRecordSet}&selected-axis=${this.selectedAxis}`).then(function (response) {
      hideLoading(_this.loadingDataTarget, elementsToToggle)
      _this.plotGraph(response.data)
      const url = '/propagation?record-set=' + _this.selectedRecordSet + `&view-option=${_this.selectedViewOption}`
      window.history.pushState(window.history.state, _this.addr, url)
    }).catch(function (e) {
      hideLoading(_this.loadingDataTarget, elementsToToggle)
      console.log(e) // todo: handle error
    })
  }

  plotGraph (csv) {
    const _this = this
    let yLabel = this.selectedRecordSet === 'votes' ? 'Time Difference (s)' : 'Delay (s)'
    let options = {
      legend: 'always',
      includeZero: true,
      legendFormatter: legendFormatter,
      labelsDiv: _this.labelsTarget,
      ylabel: yLabel,
      labelsKMB: true,
      drawPoints: true,
      strokeWidth: 0.0,
      showRangeSelector: true
    }

    if (csv.timeData) {
      let labels = ['Date', yLabel]
      let data = []
      let dataSet = []

      var extra = {
        labels: labels
      }

      let chartData = csv.timeData
      chartData.forEach(d => {
        data.push(new Date(d.date))
        data.push(d.time_difference)

        dataSet.push(data)
        data = []
      })

      _this.chartsView = new Dygraph(_this.chartsViewTarget, dataSet, { ...options, ...extra })
    } else {
      _this.chartsView = new Dygraph(_this.chartsViewTarget, csv, options)
    }
  }
}
