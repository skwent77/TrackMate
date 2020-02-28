package com.example.blockchain_v1;


import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.KeyEvent;
import android.view.Menu;
import android.view.View;
import android.view.inputmethod.EditorInfo;
import android.view.inputmethod.InputMethodManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.LinearLayout;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.drawerlayout.widget.DrawerLayout;
import androidx.navigation.Navigation;
import androidx.navigation.ui.AppBarConfiguration;

import com.google.android.material.navigation.NavigationView;
import com.google.android.material.snackbar.Snackbar;

import org.w3c.dom.Text;

public class JaeActivity extends AppCompatActivity implements TextView.OnEditorActionListener {
    LinearLayout baselayout;
    // Button button1;
    private Button submitBtn;

    private TextView text1;
    private TextView text2;
    private AppBarConfiguration mAppBarConfiguration;
    public EditText editText;
    public String forDebug;
    private String tempString;
    private boolean isSelected1;
    private boolean isSelected2;

    TextView temp;
    public String storeValueString;
    String keyword;
    @Override
//Intent를 던지는 식으로 코드를 짜야 좋으려나
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.jae_hyun);
        isSelected1=false;
        isSelected2=false;
        submitBtn = findViewById(R.id.bottom_btn);
        temp = findViewById(R.id.textViewParentTop);
        editText = findViewById(R.id.editTextFilling);
        text1 = findViewById(R.id.textViewLeftChild);
        text2 = findViewById(R.id.textViewRightChild);
        submitBtn.setOnClickListener(new View.OnClickListener() {
            public void onClick(View v) {

                text1.setOnClickListener(new View.OnClickListener() {
                    //                editText.getTex t()
                    @Override
                    public void onClick(View view) {

                        text1.setText(editText.getText());
                        //아랫줄 코드가 필요한지는 모르겠음
                        //text1쪽 누르고 나서 fin하는 경우와 editText쪽 클릭하는ㄴ 경우로 나뉨
                        editText.setOnEditorActionListener(JaeActivity.this);
                    }
                    //tempString= temp.getText().toString()++text1.getText().toString();
                    //값을 저장할거냐 표시하고 끝낼 거냐
                });

                text2.setOnClickListener(new View.OnClickListener() {

                    @Override
                    public void onClick(View view) {
                        //c짜리가 차있는 경우 -> text2누르자마자 이동하지
                        isSelected2=true;
                        changeSelectedValue();
                        //debug시에 이 블록으로 들어오는지 반드시 확인!
                        if (text2.getText().toString().length() == 0) {
                            System.out.println(1);
                        }
                        // =editText.getText().toString()
                        temp.setText(tempString);
                        text1.setText("");
                        text2.setText("");

                        editText.setText("");
                        //c자리가 안차있는 경우
                    }
                    //     +editText.getText().toString();
                });
            }
        });
//complex listener inside listener
        text1.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v){
                isSelected1=true;
                changeSelectedValue();
            }
            //     +editText.getText().toString();
        });


        }
        public void changeSelectedValue(){
        if(isSelected1){

            temp.setText(temp.getText().toString()+text1.getText().toString());
            text1.setText("");
            text2.setText("");


            editText.setText("");
        }
        if(isSelected2){

            temp.setText(temp.getText().toString()+ text2.getText().toString());
            text1.setText("");
            text2.setText("");
            editText.setText("");
        }

    }
    @Override
    public boolean onEditorAction(TextView v, int actionId, KeyEvent keyEvent) {

        switch (actionId) {
            case EditorInfo.IME_ACTION_SEARCH:

                Toast.makeText(getApplicationContext(), "검색", Toast.LENGTH_LONG).show();
                keyword = v.getText().toString();
                v.setText("");
                InputMethodManager imm = (InputMethodManager) getSystemService(getApplicationContext().INPUT_METHOD_SERVICE);
                imm.toggleSoftInputFromWindow(v.getApplicationWindowToken(), InputMethodManager.SHOW_FORCED, 0);

                break;
            default:
                return false;
        }

        return false;

    }
}


           /*
                if(text2.getText()==""){


                }
                //경우 1 a b c 중 b나 c차 있는 경우
                if(text1.getText()==""&&text2.getText()==""){



//                    temp.setText(temp.toString()+ text2.toString());
//
//                    forDebug=editText.getText().toString();
//
//                    if(text1.getText()=""&&text2.getText()==""){
//
//
//                    }else if(text2.getText()=" ") {
//
//
//                    }else{
//
//
//

                }
        });
    //    text1.setText("");
        text2.setText("");
//        text1.set(View.GONE);
        System.out.println(forDebug);
        //경우 2 a b 가 차있어 c를 입력해 b는 그대로여야
        // 경우 3  b c모두 차 있는 경우


        // Intent intent = new Intent(JaeActivity.this, JaeActivity2.class);
        //  intent.putExtra("te",temp.toString());
        //   intent.putExtra("tel",forDebug);
//                 intent.putExtra("cnum"
        Log.e("submitBtn", "debug");
        //   startActivity(intent);
        //   }

        DrawerLayout drawer = findViewById(R.id.drawer_layout);
        NavigationView navigationView = findViewById(R.id.nav_view);
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.


    }
*/
