package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	sop := `{
  "result": [
    {
      "parent": "",
      "u_asset_bar_code": "",
      "watch_list": "hitaine@yahoo-inc.com,jrex@yahoo-inc.com",
      "upon_reject": "cancel",
      "sys_updated_on": "2018-02-26 20:41:20",
      "approval_history": "",
      "skills": "",
      "u_type_of_approval": "",
      "number": "OPS0856835",
      "u_ia": "",
      "state": "1",
      "sys_created_by": "ops_install",
      "knowledge": "false",
      "order": "",
      "u_sub_status": "",
      "u_rma": "",
      "u_subcategory": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_subcategories/ef50c03c0105a500b252f7a8a2b66126",
        "value": "ef50c03c0105a500b252f7a8a2b66126"
      },
      "cmdb_ci": "",
      "impact": "3",
      "u_tile_location": "",
      "active": "true",
      "u_external_ticket_opened_date": "",
      "work_notes_list": "",
      "u_vendor": "",
      "priority": "5",
      "u_asset_serial": "",
      "u_time_assigned": "",
      "business_duration": "",
      "group_list": "",
      "u_source_from": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sc_cat_item/IA",
        "value": "IA"
      },
      "approval_set": "",
      "short_description": "Factory | gq1 | Comms-Ensemble (GLB) (United States) | for IA Request #122739",
      "correlation_display": "",
      "work_start": "",
      "additional_assignee_list": "",
      "u_requestor": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/4cbf6d7469d6b800a2c95b438b85ffbd",
        "value": "4cbf6d7469d6b800a2c95b438b85ffbd"
      },
      "sys_class_name": "u_ops_request_management",
      "closed_by": "",
      "follow_up": "",
      "u_affected_ci_id": "",
      "u_external_ticket": "",
      "reassignment_count": "0",
      "u_equipment_replaced": "",
      "u_external_ticket_closed_date": "",
      "assigned_to": "",
      "u_reopen_flag": "",
      "sla_due": "",
      "comments_and_work_notes": "",
      "u_group_approval": "",
      "u_mv_rt_json": "",
      "u_category": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_categories/c946391630f2f94857ce987d1db2d220",
        "value": "c946391630f2f94857ce987d1db2d220"
      },
      "u_reopen_count": "",
      "escalation": "0",
      "upon_approval": "proceed",
      "correlation_id": "",
      "u_asset_name": "",
      "u_total_request_duration": "",
      "u_copy_parent": "",
      "made_sla": "true",
      "u_bugzilla_sync": "",
      "u_asset_label": "",
      "u_business_service": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci_service/8434f15b0fc0cfc0abe3590be1050e86",
        "value": "8434f15b0fc0cfc0abe3590be1050e86"
      },
      "sys_updated_by": "ops_install",
      "opened_by": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/ef471b9180fc3504b824711054d609d1",
        "value": "ef471b9180fc3504b824711054d609d1"
      },
      "u_total_ownership_duration": "",
      "user_input": "",
      "sys_created_on": "2018-02-26 20:40:18",
      "u_colo_site": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci/cf1fc35bf8c01100b824e3a6d6c2214c",
        "value": "cf1fc35bf8c01100b824e3a6d6c2214c"
      },
      "sys_domain": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/global",
        "value": "global"
      },
      "u_asset_vendor": "",
      "u_last_ci_added": "",
      "u_need_by_date": "",
      "u_company_name": "Verizon Media",
      "closed_at": "",
      "business_service": "",
      "time_worked": "",
      "expected_start": "",
      "opened_at": "2018-02-26 20:40:18",
      "u_escalated": "false",
      "work_end": "",
      "u_rich_text_comments": "",
      "u_srs": "",
      "u_transaction_number": "",
      "work_notes": "",
      "u_affected_ci_ytag": "",
      "assignment_group": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/214265afed4131004ff489392e9385a7",
        "value": "214265afed4131004ff489392e9385a7"
      },
      "u_type_of_request": "general",
      "u_site": "",
      "description": "Factory Ticket:\n\nURL: https://ia.ops.yahoo.com/request/122739\n\nDetails:\n----------- Request Details ------------\n------ Host(s) ------\nType: new\nQuantity: 16\nConfiguration: SVR002615 / Custom SKU\nProLiant DL160 G664GB RAM\n1x800GB (SSD) NONErpm\n\nNew Hostname(s): iab062714[0000-0015].ostk.bm.gq1.yahoo.com\nOld Hostname(s): \nOS Information: 0 64bit Profile: 0\nRAID: No Raid\nLayer: 3\nBackplane and VLAN: pda 0\nIPv4 Type: IPv4 Private\nIPv6 Needed: No\nSite: gq1\nInstall Instructions: Stripe across at least 4 racks (ideally:  1 host per rack).  Do not need to be dedicated racks.\n\n\nHost Configuration:\n===============\n* require minimum 960GB SSD\n\n\n\n",
      "u_affected_ci": "",
      "calendar_duration": "1970-01-01 00:00:00",
      "close_notes": "",
      "u_sku": "",
      "u_ia_link": "",
      "sys_id": "452e74156fa01b00c2a9dca17b3ee4bf",
      "contact_type": "phone",
      "u_computer_room": "",
      "u_user_approval": "",
      "u_managerapproval": "false",
      "u_email_from": "",
      "urgency": "3",
      "company": {
        "link": "https://vzbuilders.service-now.com/api/now/table/core_company/f66b14e1c611227b0166c3a0df4046ff",
        "value": "f66b14e1c611227b0166c3a0df4046ff"
      },
      "u_toggle_rich_text_notes": "false",
      "activity_due": "",
      "u_last_email_update": "",
      "comments": "",
      "u_last_email_comments": "",
      "approval": "not requested",
      "due_date": "2018-03-12 20:40:18",
      "sys_mod_count": "4",
      "sys_tags": "",
      "u_approval_required": "false",
      "u_notify_ci": "false",
      "u_root_cause": "",
      "u_mv_rt_notes": "",
      "location": "",
      "u_asset_model": ""
    },
    {
      "parent": "",
      "u_asset_bar_code": "",
      "watch_list": "dcbfad7469d6b800a2c95b438b85ff30",
      "upon_reject": "cancel",
      "sys_updated_on": "2019-05-06 15:49:27",
      "approval_history": "",
      "skills": "",
      "u_type_of_approval": "AdditionalApproval",
      "number": "OPS1601587",
      "u_ia": "",
      "state": "1",
      "sys_created_by": "squarcia",
      "knowledge": "false",
      "order": "",
      "u_sub_status": "",
      "u_rma": "",
      "u_subcategory": "",
      "cmdb_ci": "",
      "impact": "3",
      "u_tile_location": "",
      "active": "true",
      "u_external_ticket_opened_date": "",
      "work_notes_list": "",
      "u_vendor": "",
      "priority": "3",
      "u_asset_serial": "",
      "u_time_assigned": "",
      "business_duration": "",
      "group_list": "",
      "u_source_from": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sc_cat_item/5298ad84ed6175044ff489392e9385f1",
        "value": "5298ad84ed6175044ff489392e9385f1"
      },
      "approval_set": "2019-05-03 16:52:21",
      "short_description": "[RE-OPEN] Move 35 Comms-Ensemble.GLB nodes from ne1 to ne1",
      "correlation_display": "",
      "work_start": "",
      "additional_assignee_list": "",
      "u_requestor": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/25cfa9b469d6b800a2c95b438b85ff32",
        "value": "25cfa9b469d6b800a2c95b438b85ff32"
      },
      "sys_class_name": "u_ops_request_management",
      "closed_by": "",
      "follow_up": "",
      "u_affected_ci_id": "",
      "u_external_ticket": "",
      "reassignment_count": "0",
      "u_equipment_replaced": "",
      "u_external_ticket_closed_date": "",
      "assigned_to": "",
      "u_reopen_flag": "yes",
      "sla_due": "",
      "comments_and_work_notes": "",
      "u_group_approval": "",
      "u_mv_rt_json": "{\"destination_colo\":\"ne1\",\"ip_type\":\"IPv4 Private\",\"need_network_req\":\"Yes\",\"u_host_in_production\":\"No\",\"u_keep_same_name\":\"Yes\"}",
      "u_category": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_categories/0d46391630f2f94857ce987d1db2d21f",
        "value": "0d46391630f2f94857ce987d1db2d21f"
      },
      "u_reopen_count": "1",
      "escalation": "0",
      "upon_approval": "proceed",
      "correlation_id": "",
      "u_asset_name": "",
      "u_total_request_duration": "",
      "u_copy_parent": "",
      "made_sla": "true",
      "u_bugzilla_sync": "",
      "u_asset_label": "",
      "u_business_service": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci_service/8434f15b0fc0cfc0abe3590be1050e86",
        "value": "8434f15b0fc0cfc0abe3590be1050e86"
      },
      "sys_updated_by": "squarcia",
      "opened_by": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/25cfa9b469d6b800a2c95b438b85ff32",
        "value": "25cfa9b469d6b800a2c95b438b85ff32"
      },
      "u_total_ownership_duration": "",
      "user_input": "",
      "sys_created_on": "2019-05-02 17:54:21",
      "u_colo_site": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci/cf1fc35bf8c01100b824e3a6d6c22159",
        "value": "cf1fc35bf8c01100b824e3a6d6c22159"
      },
      "sys_domain": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/global",
        "value": "global"
      },
      "u_asset_vendor": "",
      "u_last_ci_added": "",
      "u_need_by_date": "",
      "u_company_name": "Verizon Media",
      "closed_at": "",
      "business_service": "",
      "time_worked": "",
      "expected_start": "",
      "opened_at": "2019-05-02 17:54:20",
      "u_escalated": "false",
      "work_end": "",
      "u_rich_text_comments": "",
      "u_srs": "",
      "u_transaction_number": "",
      "work_notes": "",
      "u_affected_ci_ytag": "",
      "assignment_group": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/a94265afed4131004ff489392e9385a8",
        "value": "a94265afed4131004ff489392e9385a8"
      },
      "u_type_of_request": "move",
      "u_site": "",
      "description": "** Move request submitted by Kevin Squarcia (squarcia) on 2019-05-02 10:54:20\n\n===== Move Ticket Information ===== \n\nContact Name : Kevin Squarcia (squarcia)\nContact Details : Phone Number: +1 (408) 349-6104\r\nMobile Number: +1 (408) 890-1122\r\nYahoo ID: kevin_lavoro\r\nWork Location: US - Sunnyvale\r\nWork Country: \n\nNotes: \nPlease spread hosts across Ensemble/Omega (K8s) racks - https://docs.google.com/spreadsheets/d/1NqzaLSk0Xdx85MNhUQBimxZbk-4dkoXRUkgMDrbaDgU/edit#gid=1619474451\n\nHosts Live In Production: No\nProperty: Comms-Ensemble.GLB\nSource Colo: ne1\nDestination Colo: ne1\nIP Type: IPv4 Private\n\nKeep All Existing Name Same?: Yes\n\n Requester needs help with Network information?: Yes\n\n\n",
      "u_affected_ci": "",
      "calendar_duration": "1970-01-01 00:00:00",
      "close_notes": "The Request -OPS1601587 is closed as the Approval is Rejected",
      "u_sku": "",
      "u_ia_link": "",
      "sys_id": "8433f6a7dbc9bb40a629a026ca9619c3",
      "contact_type": "phone",
      "u_computer_room": "",
      "u_user_approval": "",
      "u_managerapproval": "false",
      "u_email_from": "",
      "urgency": "3",
      "company": {
        "link": "https://vzbuilders.service-now.com/api/now/table/core_company/6c52863ddb2e1380d602622dca96195a",
        "value": "6c52863ddb2e1380d602622dca96195a"
      },
      "u_toggle_rich_text_notes": "false",
      "activity_due": "",
      "u_last_email_update": "",
      "comments": "",
      "u_last_email_comments": "",
      "approval": "rejected",
      "due_date": "2019-05-16 17:54:20",
      "sys_mod_count": "402",
      "sys_tags": "",
      "u_approval_required": "false",
      "u_notify_ci": "false",
      "u_root_cause": "",
      "u_mv_rt_notes": "Please spread hosts across Ensemble/Omega (K8s) racks - https://docs.google.com/spreadsheets/d/1NqzaLSk0Xdx85MNhUQBimxZbk-4dkoXRUkgMDrbaDgU/edit#gid=1619474451",
      "location": "",
      "u_asset_model": ""
    },
    {
      "parent": "",
      "u_asset_bar_code": "",
      "watch_list": "hitaine@yahoo-inc.com,jrex@yahoo-inc.com",
      "upon_reject": "cancel",
      "sys_updated_on": "2018-02-26 20:41:23",
      "approval_history": "",
      "skills": "",
      "u_type_of_approval": "",
      "number": "OPS0856849",
      "u_ia": "",
      "state": "1",
      "sys_created_by": "ops_install",
      "knowledge": "false",
      "order": "",
      "u_sub_status": "",
      "u_rma": "",
      "u_subcategory": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_subcategories/ef50c03c0105a500b252f7a8a2b66126",
        "value": "ef50c03c0105a500b252f7a8a2b66126"
      },
      "cmdb_ci": "",
      "impact": "3",
      "u_tile_location": "",
      "active": "true",
      "u_external_ticket_opened_date": "",
      "work_notes_list": "",
      "u_vendor": "",
      "priority": "5",
      "u_asset_serial": "",
      "u_time_assigned": "",
      "business_duration": "",
      "group_list": "",
      "u_source_from": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sc_cat_item/IA",
        "value": "IA"
      },
      "approval_set": "",
      "short_description": "Factory | gq1 | Comms-Ensemble (GLB) (United States) | for IA Request #122739",
      "correlation_display": "",
      "work_start": "",
      "additional_assignee_list": "",
      "u_requestor": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/4cbf6d7469d6b800a2c95b438b85ffbd",
        "value": "4cbf6d7469d6b800a2c95b438b85ffbd"
      },
      "sys_class_name": "u_ops_request_management",
      "closed_by": "",
      "follow_up": "",
      "u_affected_ci_id": "",
      "u_external_ticket": "",
      "reassignment_count": "0",
      "u_equipment_replaced": "",
      "u_external_ticket_closed_date": "",
      "assigned_to": "",
      "u_reopen_flag": "",
      "sla_due": "",
      "comments_and_work_notes": "",
      "u_group_approval": "",
      "u_mv_rt_json": "",
      "u_category": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_categories/c946391630f2f94857ce987d1db2d220",
        "value": "c946391630f2f94857ce987d1db2d220"
      },
      "u_reopen_count": "",
      "escalation": "0",
      "upon_approval": "proceed",
      "correlation_id": "",
      "u_asset_name": "",
      "u_total_request_duration": "",
      "u_copy_parent": "",
      "made_sla": "true",
      "u_bugzilla_sync": "",
      "u_asset_label": "",
      "u_business_service": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci_service/8434f15b0fc0cfc0abe3590be1050e86",
        "value": "8434f15b0fc0cfc0abe3590be1050e86"
      },
      "sys_updated_by": "ops_install",
      "opened_by": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/ef471b9180fc3504b824711054d609d1",
        "value": "ef471b9180fc3504b824711054d609d1"
      },
      "u_total_ownership_duration": "",
      "user_input": "",
      "sys_created_on": "2018-02-26 20:41:07",
      "u_colo_site": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci/cf1fc35bf8c01100b824e3a6d6c2214c",
        "value": "cf1fc35bf8c01100b824e3a6d6c2214c"
      },
      "sys_domain": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/global",
        "value": "global"
      },
      "u_asset_vendor": "",
      "u_last_ci_added": "",
      "u_need_by_date": "",
      "u_company_name": "Verizon Media",
      "closed_at": "",
      "business_service": "",
      "time_worked": "",
      "expected_start": "",
      "opened_at": "2018-02-26 20:41:07",
      "u_escalated": "false",
      "work_end": "",
      "u_rich_text_comments": "",
      "u_srs": "",
      "u_transaction_number": "",
      "work_notes": "",
      "u_affected_ci_ytag": "",
      "assignment_group": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/214265afed4131004ff489392e9385a7",
        "value": "214265afed4131004ff489392e9385a7"
      },
      "u_type_of_request": "general",
      "u_site": "",
      "description": "Factory Ticket:\n\nURL: https://ia.ops.yahoo.com/request/122739\n\nDetails:\n----------- Request Details ------------\n------ Host(s) ------\nType: new\nQuantity: 16\nConfiguration: SVR002615 / Custom SKU\nProLiant DL160 G664GB RAM\n1x800GB (SSD) NONErpm\n\nNew Hostname(s): iab062714[0000-0015].ostk.bm.gq1.yahoo.com\nOld Hostname(s): \nOS Information: 0 64bit Profile: 0\nRAID: No Raid\nLayer: 3\nBackplane and VLAN: pda 0\nIPv4 Type: IPv4 Private\nIPv6 Needed: No\nSite: gq1\nInstall Instructions: Stripe across at least 4 racks (ideally:  1 host per rack).  Do not need to be dedicated racks.\n\n\nHost Configuration:\n===============\n* require minimum 960GB SSD\n\n\n\n",
      "u_affected_ci": "",
      "calendar_duration": "1970-01-01 00:00:00",
      "close_notes": "",
      "u_sku": "",
      "u_ia_link": "",
      "sys_id": "cd5e7cd96fac1b0069c88ebf2c3ee4f7",
      "contact_type": "phone",
      "u_computer_room": "",
      "u_user_approval": "",
      "u_managerapproval": "false",
      "u_email_from": "",
      "urgency": "3",
      "company": {
        "link": "https://vzbuilders.service-now.com/api/now/table/core_company/f66b14e1c611227b0166c3a0df4046ff",
        "value": "f66b14e1c611227b0166c3a0df4046ff"
      },
      "u_toggle_rich_text_notes": "false",
      "activity_due": "",
      "u_last_email_update": "",
      "comments": "",
      "u_last_email_comments": "",
      "approval": "not requested",
      "due_date": "2018-03-12 20:41:07",
      "sys_mod_count": "4",
      "sys_tags": "",
      "u_approval_required": "false",
      "u_notify_ci": "false",
      "u_root_cause": "",
      "u_mv_rt_notes": "",
      "location": "",
      "u_asset_model": ""
    },
    {
      "parent": "",
      "u_asset_bar_code": "",
      "watch_list": "",
      "upon_reject": "cancel",
      "sys_updated_on": "2019-04-23 19:09:56",
      "approval_history": "",
      "skills": "",
      "u_type_of_approval": "",
      "number": "OPS1592103",
      "u_ia": "",
      "state": "1",
      "sys_created_by": "rguo",
      "knowledge": "false",
      "order": "",
      "u_sub_status": "",
      "u_rma": "",
      "u_subcategory": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_subcategories/9446f51630f2f94857ce987d1db2d2b4",
        "value": "9446f51630f2f94857ce987d1db2d2b4"
      },
      "cmdb_ci": "",
      "impact": "3",
      "u_tile_location": "",
      "active": "true",
      "u_external_ticket_opened_date": "",
      "work_notes_list": "",
      "u_vendor": "",
      "priority": "5",
      "u_asset_serial": "",
      "u_time_assigned": "",
      "business_duration": "",
      "group_list": "",
      "u_source_from": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sc_cat_item/c033b06c282b2500a2c9e67329f493ae",
        "value": "c033b06c282b2500a2c9e67329f493ae"
      },
      "approval_set": "",
      "short_description": "HW-Repair Network for kubeetcd3.prod2.ensemble.ne1.yahoo.com",
      "correlation_display": "",
      "work_start": "",
      "additional_assignee_list": "",
      "u_requestor": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/f9ed64e06fdafa4407948ebf2c3ee484",
        "value": "f9ed64e06fdafa4407948ebf2c3ee484"
      },
      "sys_class_name": "u_ops_request_management",
      "closed_by": "",
      "follow_up": "",
      "u_affected_ci_id": "",
      "u_external_ticket": "",
      "reassignment_count": "0",
      "u_equipment_replaced": "",
      "u_external_ticket_closed_date": "",
      "assigned_to": "",
      "u_reopen_flag": "",
      "sla_due": "",
      "comments_and_work_notes": "",
      "u_group_approval": "",
      "u_mv_rt_json": "",
      "u_category": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_categories/5c46f51630f2f94857ce987d1db2d2b2",
        "value": "5c46f51630f2f94857ce987d1db2d2b2"
      },
      "u_reopen_count": "",
      "escalation": "0",
      "upon_approval": "proceed",
      "correlation_id": "",
      "u_asset_name": "",
      "u_total_request_duration": "",
      "u_copy_parent": "",
      "made_sla": "true",
      "u_bugzilla_sync": "",
      "u_asset_label": "",
      "u_business_service": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci_service/8434f15b0fc0cfc0abe3590be1050e86",
        "value": "8434f15b0fc0cfc0abe3590be1050e86"
      },
      "sys_updated_by": "rguo",
      "opened_by": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/f9ed64e06fdafa4407948ebf2c3ee484",
        "value": "f9ed64e06fdafa4407948ebf2c3ee484"
      },
      "u_total_ownership_duration": "",
      "user_input": "",
      "sys_created_on": "2019-04-23 19:09:56",
      "u_colo_site": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci/cf1fc35bf8c01100b824e3a6d6c22159",
        "value": "cf1fc35bf8c01100b824e3a6d6c22159"
      },
      "sys_domain": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/global",
        "value": "global"
      },
      "u_asset_vendor": "",
      "u_last_ci_added": "",
      "u_need_by_date": "",
      "u_company_name": "Verizon Media",
      "closed_at": "",
      "business_service": "",
      "time_worked": "",
      "expected_start": "",
      "opened_at": "2019-04-23 19:09:56",
      "u_escalated": "false",
      "work_end": "",
      "u_rich_text_comments": "",
      "u_srs": "",
      "u_transaction_number": "",
      "work_notes": "",
      "u_affected_ci_ytag": "",
      "assignment_group": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/5122e96fed4131004ff489392e938534",
        "value": "5122e96fed4131004ff489392e938534"
      },
      "u_type_of_request": "general",
      "u_site": "",
      "description": "Description of the problem: node is not pingable or reachable after reclone\nAllowed downtime: No downtime specified\nRequest part of an IA request?: No\nHost(s) in Production: Yes\nCMR ID: NA\nSeverity: Low\n",
      "u_affected_ci": "",
      "calendar_duration": "1970-01-01 00:00:00",
      "close_notes": "",
      "u_sku": "",
      "u_ia_link": "",
      "sys_id": "e95e64f8db0d77005f7773e1ba961948",
      "contact_type": "phone",
      "u_computer_room": "",
      "u_user_approval": "",
      "u_managerapproval": "false",
      "u_email_from": "",
      "urgency": "3",
      "company": {
        "link": "https://vzbuilders.service-now.com/api/now/table/core_company/6c52863ddb2e1380d602622dca96195a",
        "value": "6c52863ddb2e1380d602622dca96195a"
      },
      "u_toggle_rich_text_notes": "false",
      "activity_due": "",
      "u_last_email_update": "",
      "comments": "",
      "u_last_email_comments": "",
      "approval": "not requested",
      "due_date": "2019-05-07 19:09:56",
      "sys_mod_count": "0",
      "sys_tags": "",
      "u_approval_required": "false",
      "u_notify_ci": "false",
      "u_root_cause": "",
      "u_mv_rt_notes": "",
      "location": "",
      "u_asset_model": ""
    },
    {
      "parent": {
        "link": "https://vzbuilders.service-now.com/api/now/table/task/1e9134aadb85b740f47692c5d496193b",
        "value": "1e9134aadb85b740f47692c5d496193b"
      },
      "u_asset_bar_code": "",
      "watch_list": "",
      "upon_reject": "cancel",
      "sys_updated_on": "2019-05-02 21:20:27",
      "approval_history": "",
      "skills": "",
      "u_type_of_approval": "",
      "number": "OPS1601842",
      "u_ia": "",
      "state": "1",
      "sys_created_by": "coxswain_snow",
      "knowledge": "false",
      "order": "",
      "u_sub_status": "",
      "u_rma": "",
      "u_subcategory": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_subcategories/7846f51630f2f94857ce987d1db2d2bb",
        "value": "7846f51630f2f94857ce987d1db2d2bb"
      },
      "cmdb_ci": "",
      "impact": "3",
      "u_tile_location": "",
      "active": "true",
      "u_external_ticket_opened_date": "",
      "work_notes_list": "",
      "u_vendor": "",
      "priority": "5",
      "u_asset_serial": "",
      "u_time_assigned": "",
      "business_duration": "",
      "group_list": "",
      "u_source_from": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sc_cat_item/5298ad84ed6175044ff489392e9385f1",
        "value": "5298ad84ed6175044ff489392e9385f1"
      },
      "approval_set": "",
      "short_description": "Move 27 Comms-Ensemble.GLB nodes from ir2 to ir2 - Site Ops Child tkt",
      "correlation_display": "",
      "work_start": "",
      "additional_assignee_list": "",
      "u_requestor": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/25cfa9b469d6b800a2c95b438b85ff32",
        "value": "25cfa9b469d6b800a2c95b438b85ff32"
      },
      "sys_class_name": "u_ops_request_management",
      "closed_by": "",
      "follow_up": "",
      "u_affected_ci_id": "",
      "u_external_ticket": "",
      "reassignment_count": "0",
      "u_equipment_replaced": "",
      "u_external_ticket_closed_date": "",
      "assigned_to": "",
      "u_reopen_flag": "",
      "sla_due": "",
      "comments_and_work_notes": "",
      "u_group_approval": "",
      "u_mv_rt_json": "",
      "u_category": {
        "link": "https://vzbuilders.service-now.com/api/now/table/u_categories/8c46f51630f2f94857ce987d1db2d2b1",
        "value": "8c46f51630f2f94857ce987d1db2d2b1"
      },
      "u_reopen_count": "",
      "escalation": "0",
      "upon_approval": "proceed",
      "correlation_id": "",
      "u_asset_name": "",
      "u_total_request_duration": "",
      "u_copy_parent": "",
      "made_sla": "true",
      "u_bugzilla_sync": "",
      "u_asset_label": "",
      "u_business_service": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci_service/8434f15b0fc0cfc0abe3590be1050e86",
        "value": "8434f15b0fc0cfc0abe3590be1050e86"
      },
      "sys_updated_by": "system",
      "opened_by": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user/30b4c964300bf58857ce987d1db2d224",
        "value": "30b4c964300bf58857ce987d1db2d224"
      },
      "u_total_ownership_duration": "",
      "user_input": "",
      "sys_created_on": "2019-05-02 21:19:26",
      "u_colo_site": {
        "link": "https://vzbuilders.service-now.com/api/now/table/cmdb_ci/0f1fc35bf8c01100b824e3a6d6c2217a",
        "value": "0f1fc35bf8c01100b824e3a6d6c2217a"
      },
      "sys_domain": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/global",
        "value": "global"
      },
      "u_asset_vendor": "",
      "u_last_ci_added": "",
      "u_need_by_date": "",
      "u_company_name": "Verizon Media",
      "closed_at": "",
      "business_service": "",
      "time_worked": "",
      "expected_start": "",
      "opened_at": "2019-05-02 21:19:25",
      "u_escalated": "false",
      "work_end": "",
      "u_rich_text_comments": "",
      "u_srs": "",
      "u_transaction_number": "",
      "work_notes": "",
      "u_affected_ci_ytag": "",
      "assignment_group": {
        "link": "https://vzbuilders.service-now.com/api/now/table/sys_user_group/5122e96fed4131004ff489392e938534",
        "value": "5122e96fed4131004ff489392e938534"
      },
      "u_type_of_request": "move",
      "u_site": "",
      "description": "Ticket opened by automated Workflow system. 29536\n For questions about the Workflow system, contact workflow-admins@oath.com.\n\nTicket handling instructions:\nIF YOUR WORK IS COMPLETE:\nSet the ticket State to 'Workflow Confirm'. \nWorkflow will automatically run QA checks and close your ticket. \n \nNOTE: Please do not close this ticket yourself – Workflow will automatically reopen it.\n \n IF THERE IS A PROBLEM WITH TICKET:\n 1) For issues with the request or ticket data, contact the requestor directly.\n 2) For issues with the Workflow system, set the ticket State to 'Workflow Problem' and a workflow admin will respond to this ticket.\n\n** Move request submitted by Kevin Squarcia (squarcia) on 2019-04-28 22:55:38\n\n===== Move Ticket Information ===== \n\nContact Name : Kevin Squarcia (squarcia)\nContact Details : Phone Number: +1 (408) 349-6104\r\nMobile Number: +1 (408) 890-1122\r\nYahoo ID: kevin_lavoro\r\nWork Location: US - Sunnyvale\r\nWork Country: \n\nNotes: \nneed to move these hosts to Ensemble racks - https://docs.google.com/spreadsheets/d/1NqzaLSk0Xdx85MNhUQBimxZbk-4dkoXRUkgMDrbaDgU/edit#gid=1013288024\n\nHosts Live In Production: No\nProperty: Comms-Ensemble.GLB\nSource Colo: ir2\nDestination Colo: ir2\nIP Type: IPv4 Private\n\nKeep All Existing Name Same?: Yes\n\n Requester needs help with Network information?: Yes\n\n\n",
      "u_affected_ci": "",
      "calendar_duration": "1970-01-01 00:00:00",
      "close_notes": "",
      "u_sku": "",
      "u_ia_link": "",
      "sys_id": "f2b22f27db09f380f47692c5d49619fb",
      "contact_type": "phone",
      "u_computer_room": "",
      "u_user_approval": "",
      "u_managerapproval": "false",
      "u_email_from": "",
      "urgency": "3",
      "company": {
        "link": "https://vzbuilders.service-now.com/api/now/table/core_company/f66b14e1c611227b0166c3a0df4046ff",
        "value": "f66b14e1c611227b0166c3a0df4046ff"
      },
      "u_toggle_rich_text_notes": "false",
      "activity_due": "",
      "u_last_email_update": "",
      "comments": "",
      "u_last_email_comments": "",
      "approval": "not requested",
      "due_date": "2019-05-16 21:19:25",
      "sys_mod_count": "3",
      "sys_tags": "",
      "u_approval_required": "false",
      "u_notify_ci": "true",
      "u_root_cause": "",
      "u_mv_rt_notes": "need to move these hosts to Ensemble racks - https://docs.google.com/spreadsheets/d/1NqzaLSk0Xdx85MNhUQBimxZbk-4dkoXRUkgMDrbaDgU/edit#gid=1013288024",
      "location": "",
      "u_asset_model": ""
    }
  ]
}
`
	type snowlinkobj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}

	type ServiceNowResponse struct {
		Parent                    string      `json:"parent"`
		UAssetBarCode             string      `json:"u_asset_bar_code"`
		WatchList                 string      `json:"watch_list"`
		UponReject                string      `json:"upon_reject"`
		SysUpdatedOn              string      `json:"sys_updated_on"`
		ApprovalHistory           string      `json:"approval_history"`
		Skills                    string      `json:"skills"`
		UTypeOfApproval           string      `json:"u_type_of_approval"`
		Number                    string      `json:"number"`
		UIa                       string      `json:"u_ia"`
		State                     string      `json:"state"`
		SysCreatedBy              string      `json:"sys_created_by"`
		Knowledge                 string      `json:"knowledge"`
		Order                     string      `json:"order"`
		USubStatus                string      `json:"u_sub_status"`
		URma                      string      `json:"u_rma"`
		USubcategory              snowlinkobj `json:"u_subcategory"`
		CmdbCi                    string      `json:"cmdb_ci"`
		Impact                    string      `json:"impact"`
		UTileLocation             string      `json:"u_tile_location"`
		Active                    string      `json:"active"`
		UExternalTicketOpenedDate string      `json:"u_external_ticket_opened_date"`
		WorkNotesList             string      `json:"work_notes_list"`
		UVendor                   string      `json:"u_vendor"`
		Priority                  string      `json:"priority"`
		UAssetSerial              string      `json:"u_asset_serial"`
		UTimeAssigned             string      `json:"u_time_assigned"`
		BusinessDuration          string      `json:"business_duration"`
		GroupList                 string      `json:"group_list"`
		USourceFrom               snowlinkobj `json:"u_source_from"`
		ApprovalSet               string      `json:"approval_set"`
		ShortDescription          string      `json:"short_description"`
		CorrelationDisplay        string      `json:"correlation_display"`
		WorkStart                 string      `json:"work_start"`
		AdditionalAssigneeList    string      `json:"additional_assignee_list"`
		URequestor                snowlinkobj `json:"u_requestor"`
		SysClassName              string      `json:"sys_class_name"`
		ClosedBy                  string      `json:"closed_by"`
		FollowUp                  string      `json:"follow_up"`
		UAffectedCiID             string      `json:"u_affected_ci_id"`
		UExternalTicket           string      `json:"u_external_ticket"`
		ReassignmentCount         string      `json:"reassignment_count"`
		UEquipmentReplaced        string      `json:"u_equipment_replaced"`
		UExternalTicketClosedDate string      `json:"u_external_ticket_closed_date"`
		AssignedTo                string      `json:"assigned_to"`
		UReopenFlag               string      `json:"u_reopen_flag"`
		SLADue                    string      `json:"sla_due"`
		CommentsAndWorkNotes      string      `json:"comments_and_work_notes"`
		UGroupApproval            string      `json:"u_group_approval"`
		UMvRtJSON                 string      `json:"u_mv_rt_json"`
		UCategory                 snowlinkobj `json:"u_category"`
		UReopenCount              string      `json:"u_reopen_count"`
		Escalation                string      `json:"escalation"`
		UponApproval              string      `json:"upon_approval"`
		CorrelationID             string      `json:"correlation_id"`
		UAssetName                string      `json:"u_asset_name"`
		UTotalRequestDuration     string      `json:"u_total_request_duration"`
		UCopyParent               string      `json:"u_copy_parent"`
		MadeSLA                   string      `json:"made_sla"`
		UBugzillaSync             string      `json:"u_bugzilla_sync"`
		UAssetLabel               string      `json:"u_asset_label"`
		UBusinessService          snowlinkobj `json:"u_business_service"`
		SysUpdatedBy              string      `json:"sys_updated_by"`
		OpenedBy                  snowlinkobj `json:"opened_by"`
		UTotalOwnershipDuration   string      `json:"u_total_ownership_duration"`
		UserInput                 string      `json:"user_input"`
		SysCreatedOn              string      `json:"sys_created_on"`
		UColoSite                 snowlinkobj `json:"u_colo_site"`
		SysDomain                 snowlinkobj `json:"sys_domain"`
		UAssetVendor              string      `json:"u_asset_vendor"`
		ULastCiAdded              string      `json:"u_last_ci_added"`
		UNeedByDate               string      `json:"u_need_by_date"`
		UCompanyName              string      `json:"u_company_name"`
		ClosedAt                  string      `json:"closed_at"`
		BusinessService           string      `json:"business_service"`
		TimeWorked                string      `json:"time_worked"`
		ExpectedStart             string      `json:"expected_start"`
		OpenedAt                  string      `json:"opened_at"`
		UEscalated                string      `json:"u_escalated"`
		WorkEnd                   string      `json:"work_end"`
		URichTextComments         string      `json:"u_rich_text_comments"`
		USrs                      string      `json:"u_srs"`
		UTransactionNumber        string      `json:"u_transaction_number"`
		WorkNotes                 string      `json:"work_notes"`
		UAffectedCiYtag           string      `json:"u_affected_ci_ytag"`
		AssignmentGroup           snowlinkobj `json:"assignment_group"`
		UTypeOfRequest            string      `json:"u_type_of_request"`
		USite                     string      `json:"u_site"`
		Description               string      `json:"description"`
		UAffectedCi               string      `json:"u_affected_ci"`
		CalendarDuration          string      `json:"calendar_duration"`
		CloseNotes                string      `json:"close_notes"`
		USku                      string      `json:"u_sku"`
		UIaLink                   string      `json:"u_ia_link"`
		SysID                     string      `json:"sys_id"`
		ContactType               string      `json:"contact_type"`
		UComputerRoom             string      `json:"u_computer_room"`
		UUserApproval             string      `json:"u_user_approval"`
		UManagerapproval          string      `json:"u_managerapproval"`
		UEmailFrom                string      `json:"u_email_from"`
		Urgency                   string      `json:"urgency"`
		Company                   snowlinkobj `json:"company"`
		UToggleRichTextNotes      string      `json:"u_toggle_rich_text_notes"`
		ActivityDue               string      `json:"activity_due"`
		ULastEmailUpdate          string      `json:"u_last_email_update"`
		Comments                  string      `json:"comments"`
		ULastEmailComments        string      `json:"u_last_email_comments"`
		Approval                  string      `json:"approval"`
		DueDate                   string      `json:"due_date"`
		SysModCount               string      `json:"sys_mod_count"`
		SysTags                   string      `json:"sys_tags"`
		UApprovalRequired         string      `json:"u_approval_required"`
		UNotifyCi                 string      `json:"u_notify_ci"`
		URootCause                string      `json:"u_root_cause"`
		UMvRtNotes                string      `json:"u_mv_rt_notes"`
		Location                  string      `json:"location"`
		UAssetModel               string      `json:"u_asset_model"`
	}

	//SnowResultArray when you have multiple responses
	type SnowResultArray struct {
		Result []ServiceNowResponse
	}

	var sobj SnowResultArray
	err := json.Unmarshal([]byte(sop), sobj)
	if err != nil {
		fmt.Printf("Did not work %s\n", err.String())

	}

}
