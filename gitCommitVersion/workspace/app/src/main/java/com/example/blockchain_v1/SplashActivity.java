package com.example.blockchain_v1;
//
//import androidx.appcompat.app.AppCompatActivity;
//
//public class SplashActivity extends AppCompatActivity {
//
//
//}

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;

import androidx.appcompat.app.AppCompatActivity;


public class SplashActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_splash);
        Intent intent=new Intent(this,RealActivity.class);
        startActivity(  intent);
        Handler hd = new Handler();
        hd.postDelayed(new SplashHandler(), 500); // 1초 후에 hd handler 실행  3000ms = 3초
//delay시간이 필요하거나 한경우 바로 어떤 명령을 실행하지 않고 딜레이 갖고 싶을 때
    }

    private class SplashHandler implements Runnable {
        public void run() {
            startActivity(new Intent(getApplication(), MainActivity.class)); //로딩이 끝난 후, ChoiceFunction 이동
            SplashActivity.this.finish(); // 로딩페이지 Activity stack에서 제거
        }
    }

    @Override
    public void onBackPressed() {
        //초반 플래시 화면에서 넘어갈때 뒤로가기 버튼 못누르게 함
    }

}