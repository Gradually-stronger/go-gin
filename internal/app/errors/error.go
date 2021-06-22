package errors

import (
	"github.com/pkg/errors"
)

// 定义错误函数的别名
var (
	New       = errors.New
	Wrap      = errors.Wrap
	Wrapf     = errors.Wrapf
	WithStack = errors.WithStack
)

// 定义错误
var (
	// 公共错误
	ErrNotFound                = New("资源不存在或已被删除")
	ErrMethodNotAllow          = New("方法不被允许")
	ErrBadRequest              = New("请求发生错误")
	ErrParameterNotEnough      = New("请求参数不足")
	ErrInvalidRequestParameter = New("无效的请求参数")
	ErrFewRequestParameter     = New("缺少请求参数")
	ErrTooManyRequests         = New("请求过于频繁")
	ErrUnknownQuery            = New("未知的查询类型")
	ErrInvalidParent           = New("无效的父级节点")
	ErrNotAllowDeleteWithChild = New("含有子级，不能删除")
	ErrResourceExists          = New("资源已经存在")
	ErrResourceNotAllowDelete  = New("资源不允许删除")
	ErrAlreadyDone             = New("操作已完成，请勿重复操作")
	ErrNameDuplicate           = New("名称重复")

	//角色管理
	ErrRoleNameExists = New("角色名重复，无法添加")

	// 权限错误
	ErrNoPerm         = New("无访问权限")
	ErrNoResourcePerm = New("无资源的访问权限")

	// 用户错误
	ErrInvalidUserName = New("无效的用户名")
	ErrInvalidPassword = New("无效的密码")
	ErrInvalidUser     = New("无效的用户")
	ErrUserDisable     = New("用户被禁用")
	ErrUserNotEmptyPwd = New("密码不允许为空")

	// login
	ErrLoginNotAllowModifyPwd = New("不允许修改密码")
	ErrLoginInvalidOldPwd     = New("旧密码不正确")
	ErrLoginInvalidVerifyCode = New("无效的验证码")

	// 成本核算
	ErrNoProjCostItem  = New("缺少成本测算")
	ErrNoProjSalesPlan = New("缺少销售计划")
	ErrNoTaxIncome     = New("未设置所得税")
	ErrNoTaxStamp      = New("未设置印花税")
	ErrNoTaxUse        = New("未设置使用税")
	ErrNoTaxContract   = New("未设置契税")
	ErrNoTaxAdditional = New("未设置地方附加税")
	ErrNoTaxOutput     = New("未设置增值税销项税")

	//合同管理
	ErrNotRightStatusForSetSN        = New("当前状态不能添加编号")
	ErrNotRightStatusForCancelCommit = New("该合同状态不为审核中，不能进行撤销审核操作")
	ErrComContractNotPassCheck       = New("该合同还未通过审核")
	ErrComContractNotCommit          = New("该合同还未提交审核")
	ErrNoComContractSignDate         = New("请输入合同签署日期")
	ErrNoComContractSN               = New("请输入合同正式编号")
	ErrNoSettlement                  = New("该合同信息没有选择结算")
	ErrNoComContract                 = New("没有找到对应的合同")
	// 合约规划
	ErrNoInCome        = New("未生成目标成本")
	ErrReferStasusDone = New("引用已结束")
	// 建筑管理
	ErrNoChildrenBuilding     = New("拆分的建筑缺少子建筑")
	ErrRoomCantSplit          = New("门牌无法继续拆分")
	ErrNameExists             = New("建筑名已存在")
	ErrNameUnExists           = New("建筑名称有误")
	ErrBuildingIsRenting      = New("建筑租赁中，禁止操作")
	ErrDataRelationship       = New("数据关系错误")
	ErrNotMinimumBuild        = New("查询的建筑不是最小可操作建筑")
	ErrAreaCantBeZero         = New("最小单元面积不能为0")
	ErrBadUsageType           = New("错误的建筑规划用途")
	ErrBuildingInProjectGroup = New("建筑存在于项目组中，禁止操作")
	ErrBuildingInAssetLease   = New("建筑关联租赁决策文件，禁止操作")
	ErrBuildingInRentContract = New("建筑关联合同，禁止操作")

	// 资产项目管理
	ErrHasBuilding = New("项目下有建筑，禁止删除")
	ErrHasGroup    = New("项目下有项目组，禁止删除")
	ErrOrgErr      = New("查询无关联组织")
	ErrComErr      = New("查询无关联公司")

	// 资产项目组管理
	ErrFewGroupManager = New("缺少项目组负责人")
	ErrDuplicateRules  = New("建筑规则与现有数据重复")
	ErrHasContract     = New("项目组有关联合同，禁止删除")
	ErrHasUser         = New("项目组包含用户，禁止删除")

	// 租赁文件管理
	ErrUpdateStatus = New("租赁文件更新状态错误")

	// 合约规划审核
	ErrorUncommitPlan = New("合约规划未提交审核")

	//租赁合同相关错误
	ErrorRentContractDenyCantEdit = New("合同状态不允许编辑")
	ErrorRentContractEditOthers   = New("您不能编辑别人的合同信息")
	ErrorRentContractFilesErr     = New("合同附件pdf格式错误")
	ErrorRentContractRenewal      = New("当前合同不允许执行续租操作")
	ErrorRentContractNoBuilding   = New("合同没有房间")

	// 租户相关
	ErrorInvalidPhoneNumber = New("电话号码格式不正确")
	ErrorTenantType         = New("租户类型有误")

	// 员工档案管理
	ErrorEmployeeIDNumber    = New("身份证号码格式不正确")
	ErrorEmployeePhoneNumber = New("手机号码格式不正确")
	ErrorEmployeeEmail       = New("电子邮箱格式不正确")

	// 人力资源模板
	ErrorTemplateUsing   = New("有公司正在使用此模板，不能删除")
	ErrArchiveNotFound   = New("档案不存在")
	ErrArchiveNotMatched = New("无匹配档案")
	ErrManyArchive       = New("多条匹配档案")
	ErrArchiveMatched    = New("档案已匹配帐号，请登录")
	ErrCompanyNotFound   = New("查询的公司不存在")
	ErrRoleNotFound      = New("指定公司下未找到\"个人档案管理\"角色")
	ErrNotMatchedArchive = New("用户未关联档案")
	ErrNotJoinCompany    = New("用户未关联公司")

	// 营销 当前合同已存在内容变更
	ErrorHasChange   = New("当前合同已存在内容变更，请处理")
	ErrorPlaceOnFile = New("当前合同已归档，请勿重复操作")
	ErrorAreaState   = New("该建筑无预测、实售面积")

	// 资金
	ErrExistPayingApplied = New("当前合同有一笔资金申请正在处理，请等待")

	// 考核填报
	ErrEvaluationSettingRepeated = New("该年度考核设定已经存在，请勿重复添加")
	ErrEvaluationNoSetting       = New("该年度考核缺少填报设定信息，请联系管理员")
	ErrEvaluationEditTimeErr     = New("考核设定中的填报时间格式有误")
	ErrEvaluationSubmitTimeout   = New("已经超过提交时间，如有必要请联系管理员")
	ErrEvaluationSubmitFail      = New("提交审核失败")
	ErrEvaluationSubmitted       = New("本年度考核报表已经提交，请勿重复操作")
	ErrEvaluationPower           = New("请确保每个季度的权重都为100.00")

	//计划管理
	//ErrorExecutionPercentage = New("当前进度请大于制定当前进度")
	//ErrorUnderApproval       = New("审批中请勿再次提交")
	//ErrorApproved            = New("审批以通过请勿再次提交")

	ErrorFollowUpImpact          = New("请填写对后续工作影响")
	ErrorSolution                = New("请填写解决方案")
	ErrorEstimatedCompletionTime = New("请填写预计完成时间")
	// 交易管理-预约排号
	ErrHomeInfo = New("房间信息有误")
)

func init() {
	// 公共错误
	newBadRequestError(ErrBadRequest)
	newBadRequestError(ErrInvalidRequestParameter)
	newBadRequestError(ErrFewRequestParameter)
	newBadRequestError(ErrParameterNotEnough)
	newErrorCode(ErrNotFound, 404, ErrNotFound.Error(), 404)
	newErrorCode(ErrMethodNotAllow, 405, ErrMethodNotAllow.Error(), 405)
	newErrorCode(ErrTooManyRequests, 429, ErrTooManyRequests.Error(), 429)
	newBadRequestError(ErrUnknownQuery)
	newBadRequestError(ErrInvalidParent)
	newBadRequestError(ErrNotAllowDeleteWithChild)
	newBadRequestError(ErrResourceExists)
	newBadRequestError(ErrResourceNotAllowDelete)
	newBadRequestError(ErrAlreadyDone)

	newBadRequestError(ErrRoleNameExists)

	newBadRequestError(ErrNameDuplicate)

	// 权限错误
	newErrorCode(ErrNoPerm, 9999, ErrNoPerm.Error(), 401)
	newErrorCode(ErrNoResourcePerm, 401, ErrNoResourcePerm.Error(), 401)

	// 用户错误
	newBadRequestError(ErrInvalidUserName)
	newBadRequestError(ErrInvalidPassword)
	newBadRequestError(ErrInvalidUser)
	newBadRequestError(ErrUserDisable)
	newBadRequestError(ErrUserNotEmptyPwd)

	// login
	newBadRequestError(ErrLoginNotAllowModifyPwd)
	newBadRequestError(ErrLoginInvalidOldPwd)
	newBadRequestError(ErrLoginInvalidVerifyCode)

	//成本核算
	newBadRequestError(ErrNoProjCostItem)
	newBadRequestError(ErrNoProjSalesPlan)
	newBadRequestError(ErrNoTaxIncome)
	newBadRequestError(ErrNoTaxStamp)
	newBadRequestError(ErrNoTaxUse)
	newBadRequestError(ErrNoTaxContract)
	newBadRequestError(ErrNoTaxAdditional)
	newBadRequestError(ErrNoTaxOutput)

	//合同管理
	newBadRequestError(ErrNotRightStatusForSetSN)
	newBadRequestError(ErrNotRightStatusForCancelCommit)
	newBadRequestError(ErrNoSettlement)
	newBadRequestError(ErrNoComContract)

	// 合约规划
	newBadRequestError(ErrNoInCome)
	newBadRequestError(ErrReferStasusDone)

	// 建筑管理
	newBadRequestError(ErrNoChildrenBuilding)
	newBadRequestError(ErrRoomCantSplit)
	newBadRequestError(ErrNameExists)
	newBadRequestError(ErrNameUnExists)
	newBadRequestError(ErrBuildingIsRenting)
	newInternalServerError(ErrDataRelationship)
	newInternalServerError(ErrBadUsageType)
	newBadRequestError(ErrAreaCantBeZero)
	newBadRequestError(ErrNotMinimumBuild)
	newBadRequestError(ErrBuildingInProjectGroup)
	newBadRequestError(ErrBuildingInAssetLease)
	newBadRequestError(ErrBuildingInRentContract)

	// 资产项目管理
	newBadRequestError(ErrHasBuilding)
	newBadRequestError(ErrHasGroup)
	newBadRequestError(ErrOrgErr)
	newBadRequestError(ErrComErr)

	// 资产项目组管理
	newBadRequestError(ErrFewGroupManager)
	newBadRequestError(ErrDuplicateRules)
	newBadRequestError(ErrHasContract)
	newBadRequestError(ErrHasUser)

	// 租赁文件管理
	newBadRequestError(ErrUpdateStatus)

	// 合约规划审核
	newBadRequestError(ErrorUncommitPlan)

	// 租赁合同模板
	newInternalServerError(ErrorRentContractFilesErr)
	// 禁止合同续租
	newBadRequestError(ErrorRentContractRenewal)
	// 禁止编辑他人合同信息
	newBadRequestError(ErrorRentContractEditOthers)
	// 禁止编辑合同
	newBadRequestError(ErrorRentContractDenyCantEdit)
	// 当前合同没有查到关联的房间
	newBadRequestError(ErrorRentContractNoBuilding)

	// 租户管理
	newBadRequestError(ErrorInvalidPhoneNumber)
	newBadRequestError(ErrorTenantType)

	// 员工档案管理
	newBadRequestError(ErrorEmployeeIDNumber)
	newBadRequestError(ErrorEmployeeEmail)
	newBadRequestError(ErrorEmployeePhoneNumber)

	// 人力资源模板
	newBadRequestError(ErrorTemplateUsing)
	newBadRequestError(ErrArchiveNotFound)
	newBadRequestError(ErrArchiveNotMatched)
	newBadRequestError(ErrManyArchive)
	newBadRequestError(ErrArchiveMatched)
	newBadRequestError(ErrCompanyNotFound)
	newBadRequestError(ErrRoleNotFound)
	newBadRequestError(ErrNotMatchedArchive)
	newBadRequestError(ErrNotJoinCompany)

	// 营销
	newBadRequestError(ErrorHasChange)

	// 资金
	newBadRequestError(ErrExistPayingApplied)

	// 考核填报
	newBadRequestError(ErrEvaluationSettingRepeated)
	newBadRequestError(ErrEvaluationNoSetting)
	newBadRequestError(ErrEvaluationEditTimeErr)
	newBadRequestError(ErrEvaluationSubmitTimeout)
	newBadRequestError(ErrEvaluationSubmitFail)
	newBadRequestError(ErrEvaluationSubmitted)
	newBadRequestError(ErrEvaluationPower)
	//计划管理
	newBadRequestError(ErrorFollowUpImpact)
	newBadRequestError(ErrorSolution)
	newBadRequestError(ErrorEstimatedCompletionTime)

	// 交易管理-预约排号
	newBadRequestError(ErrHomeInfo)
}
