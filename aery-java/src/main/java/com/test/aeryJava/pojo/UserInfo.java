package com.test.aeryJava.pojo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDate;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserInfo {
    private Integer userId;
    private Integer userInfoId;
    private String nickName;
    private String account;
    private Integer sex;
    private LocalDate birth;
    private Integer age;
    private String email;
    private String phone;
    private String brief;
    private LocalDate createDate;
    private String avatar;
    private String currentAddress;
    private String address;
    private String info;
}
