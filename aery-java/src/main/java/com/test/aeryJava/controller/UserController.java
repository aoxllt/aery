package com.test.aeryJava.controller;

import com.test.aeryJava.pojo.Result;
import com.test.aeryJava.service.UserService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/user")
public class UserController {
    private static final Logger log = LoggerFactory.getLogger(UserController.class);
    @Autowired
    UserService userService;

    @GetMapping("/{userId}")
    public Result getUser(@PathVariable Integer userId) {
        log.info("get user id:{}", userId);
        return Result.success(userService.getUserInfo(userId));
    }

}
