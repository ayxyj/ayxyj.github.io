package main

import (
	"fmt"
	"time"
)

func main() {
	CallTest()
}

func CallTest() {
	//soli 首先获取信息 封装起来
	cal := &call{
		"主任医师",
		"内科",
		"患者基本信息",
		"急诊科",
		"读取",
		"医疗",
		"12:00",
	}
	//false
	cal1 := &call{
		"主任医师",
		"内科",
		"患者基本信息",
		"急诊科",
		"读取",
		"医疗",
		"08:00",
	}
	//false
	cal2 := &call{
		"主任医师",
		"内科",
		"医嘱",
		"急诊科",
		"读取",
		"医疗",
		"09:00",
	}
	fmt.Println("---Access Control Result---")
	fmt.Println(test111(*cal))
	fmt.Println(test111(*cal1))
	fmt.Println(test111(*cal2))
}

type call struct {
	SA_Role    string
	SA_Classes string
	RA_Type    string
	RA_Classes string
	OA_Type    string
	OA_Purpose string
	EA_Time    string
}

//msg.address
func test111(cal call) string {
	if proc01(cal) {
		return "PERMIT"
	} else {
		return "REFUSE"
	}
}

func proc02(cal call) bool {
	//主任医师，副主任医师，主治医生
	//住院医生，主任护师，副主任护师，主管护师，护士，研究机构，患者
	if cal.SA_Role == "主任医师" || cal.SA_Role == "副主任医师" {
		ZhuRenAndFuZhuRen(cal)
	}
	if cal.SA_Role == "主治医师" {
		ZhuZhiYiShi(cal)
	}
	if cal.SA_Role == "住院医生" {
		ZhuYuanYiShi(cal)
	}
	if cal.SA_Role == "主任护师" || cal.SA_Role == "副主任护师" || cal.SA_Role == "主管护师" || cal.SA_Role == "护士" {
		OtherStaffs(cal)
	}
	if cal.SA_Role == "研究机构" {
		Research(cal)
	}
	if cal.SA_Role == "患者" {
		flag := HuanZhe(cal)
		return flag
	}
}

func ZhuRenAndFuZhuRen(cal call) bool {
	//主任医师，副主任医师
	myDate, _ := time.Parse("15:04", cal.EA_Time)
	fmt.Println(myDate)
	fmt.Println(myDate.Format("15:00"))
	if myDate.Format("15:00") >= "09:00" && myDate.Format("15:00") <= "17:00" {

	} else {
		fmt.Println("下班时间！")
		return false
	}
}

func ZhuZhiYiShi(cal call) bool {
	//主治医生
	myDate, _ := time.Parse("15:04", cal.EA_Time)
	fmt.Println(myDate)
	fmt.Println(myDate.Format("15:00"))
	if myDate.Format("15:00") >= "09:00" && myDate.Format("15:00") <= "17:00" {

	} else {
		fmt.Println("下班时间！")
		return false
	}
}

func ZhuYuanYiShi(cal call) {
	//住院医生
}
func OtherStaffs(cal call) {
	//主任护师，副主任护师，主管护师，护士
}
func Research(cal call) bool {
	//{研究机构}
	//{无}
	//{门诊病历 住院病历 护理病历 医嘱}
	if Judge_RA_TYPE(cal) {
		//{急诊科 内科 外科 妇产科 预防保健科 药房 化验科 X光室 消毒供应室}
		if Judge_RA_CLASSES(cal){
			//{读取 打印}
			if cal.OA_Type=="读取" || cal.OA_Type=="打印"{
				//{研究 教学}
				if cal.OA_Purpose=="研究" || cal.OA_Purpose=="教学"{
					return true
				}else{
					fmt.Println("操作目的错误！")
					return false
				}
			}else{
				fmt.Println("操作类型错误！")
				return false
			}
		}else{
			return false
		}
	} else {
		return false
	}
}

func HuanZhe(cal call) bool {
	//患者
	//患者基本信息 门诊病历 住院病历 护理病历 医嘱
	if cal.RA_Type == "患者基本信息" || Judge_RA_TYPE(cal){
		//急诊科 内科 外科 妇产科 预防保健科 药房 化验科 X光室 消毒供应室
		if Judge_RA_CLASSES(cal) {
			//读取 打印
			if cal.OA_Type == "读取" || cal.OA_Type == "打印" {
				return true
			} else {
				fmt.Println("操作不当！")
				return false
			}
		} else {
			fmt.Println("资源所属科室错误！")
			return false
		}
	} else {
		fmt.Println("资源类型错误！")
		return false
	}
}
func Judge_RA_TYPE(cal call) bool{
	//{门诊病历 住院病历 护理病历 医嘱}
	if cal.RA_Type == "门诊病历" || cal.RA_Type == "住院病历" ||
		cal.RA_Type == "护理病历" || cal.RA_Type == "医嘱" {
		return true
	} else {
		fmt.Println("资源类型错误！")
		return false
	}
}
func Judge_RA_CLASSES(cal call) bool {
	if cal.RA_Classes == "急诊科" || cal.RA_Classes == "内科" || cal.RA_Classes == "外科" || cal.RA_Classes == "妇产科" ||
		cal.RA_Classes == "预防保健科" || cal.RA_Classes == "药房" || cal.RA_Classes == "化验科" || cal.RA_Classes == "X光室" ||
		cal.RA_Classes == "消毒供应室" {
		return true
	} else {
		fmt.Println("资源所属科室错误！")
		return false
	}
}

func proc01(cal call) bool {
	//时间处理
	myDate, _ := time.Parse("15:04", cal.EA_Time)
	if myDate.Format("15:00") >= "09:00" && myDate.Format("15:00") <= "17:00" {
		if cal.SA_Role == "主任医师" || cal.SA_Role == "副主任医师" {
			if cal.SA_Classes == "急诊科" || cal.SA_Classes == "内科" {
				if cal.RA_Type == "门诊病历" || cal.RA_Type == "住院病历" || cal.RA_Type == "患者基本信息" {
					if cal.RA_Classes == "急诊科" || cal.RA_Classes == "内科" || cal.RA_Classes == "外科" {
						if cal.OA_Type == "读取" || cal.OA_Type == "写入" {
							if cal.OA_Purpose == "医疗" || cal.OA_Purpose == "研究" {
								return true
							} else {
								//log print or error statement
								fmt.Println("目的不纯！")
								return false
							}
						} else {
							fmt.Println("操作不当！")
							return false
						}
					} else {
						fmt.Println("资源所属科室错误！")
						return false
					}
				} else {
					fmt.Println("资源类型错误！")
					return false
				}
			} else {
				fmt.Println("访问者科室错误！")
				return false
			}
		} else {
			fmt.Println("访问者角色错误！")
			return false
		}
	} else {
		fmt.Println("下班时间！")
		return false
	}
}
