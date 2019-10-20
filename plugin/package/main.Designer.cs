namespace package
{
    partial class FormMain
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(FormMain));
            this.tab = new System.Windows.Forms.TabControl();
            this.tabPage1 = new System.Windows.Forms.TabPage();
            this.comboBox_pack = new System.Windows.Forms.ComboBox();
            this.label4 = new System.Windows.Forms.Label();
            this.button_pack_execute = new System.Windows.Forms.Button();
            this.progressBar_pack = new System.Windows.Forms.ProgressBar();
            this.button_pack_add = new System.Windows.Forms.Button();
            this.button_pack_delete = new System.Windows.Forms.Button();
            this.label3 = new System.Windows.Forms.Label();
            this.button_pack_select = new System.Windows.Forms.Button();
            this.textBox_pack = new System.Windows.Forms.TextBox();
            this.label2 = new System.Windows.Forms.Label();
            this.listView_pack = new System.Windows.Forms.ListView();
            this.label1 = new System.Windows.Forms.Label();
            this.tabPage2 = new System.Windows.Forms.TabPage();
            this.tab.SuspendLayout();
            this.tabPage1.SuspendLayout();
            this.SuspendLayout();
            // 
            // tab
            // 
            this.tab.Controls.Add(this.tabPage1);
            this.tab.Controls.Add(this.tabPage2);
            this.tab.Font = new System.Drawing.Font("Kristen ITC", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.tab.Location = new System.Drawing.Point(12, 12);
            this.tab.Name = "tab";
            this.tab.SelectedIndex = 0;
            this.tab.Size = new System.Drawing.Size(760, 537);
            this.tab.TabIndex = 0;
            // 
            // tabPage1
            // 
            this.tabPage1.Controls.Add(this.comboBox_pack);
            this.tabPage1.Controls.Add(this.label4);
            this.tabPage1.Controls.Add(this.button_pack_execute);
            this.tabPage1.Controls.Add(this.progressBar_pack);
            this.tabPage1.Controls.Add(this.button_pack_add);
            this.tabPage1.Controls.Add(this.button_pack_delete);
            this.tabPage1.Controls.Add(this.label3);
            this.tabPage1.Controls.Add(this.button_pack_select);
            this.tabPage1.Controls.Add(this.textBox_pack);
            this.tabPage1.Controls.Add(this.label2);
            this.tabPage1.Controls.Add(this.listView_pack);
            this.tabPage1.Controls.Add(this.label1);
            this.tabPage1.Location = new System.Drawing.Point(4, 26);
            this.tabPage1.Name = "tabPage1";
            this.tabPage1.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage1.Size = new System.Drawing.Size(752, 507);
            this.tabPage1.TabIndex = 0;
            this.tabPage1.Text = "pack";
            this.tabPage1.UseVisualStyleBackColor = true;
            // 
            // comboBox_pack
            // 
            this.comboBox_pack.FormattingEnabled = true;
            this.comboBox_pack.Location = new System.Drawing.Point(33, 389);
            this.comboBox_pack.Name = "comboBox_pack";
            this.comboBox_pack.Size = new System.Drawing.Size(607, 25);
            this.comboBox_pack.TabIndex = 12;
            // 
            // label4
            // 
            this.label4.AutoSize = true;
            this.label4.Location = new System.Drawing.Point(30, 369);
            this.label4.Name = "label4";
            this.label4.Size = new System.Drawing.Size(69, 17);
            this.label4.TabIndex = 10;
            this.label4.Text = "pack type:";
            // 
            // button_pack_execute
            // 
            this.button_pack_execute.Location = new System.Drawing.Point(646, 451);
            this.button_pack_execute.Name = "button_pack_execute";
            this.button_pack_execute.Size = new System.Drawing.Size(75, 23);
            this.button_pack_execute.TabIndex = 9;
            this.button_pack_execute.Text = "execute";
            this.button_pack_execute.UseVisualStyleBackColor = true;
            // 
            // progressBar_pack
            // 
            this.progressBar_pack.Location = new System.Drawing.Point(33, 451);
            this.progressBar_pack.Name = "progressBar_pack";
            this.progressBar_pack.Size = new System.Drawing.Size(607, 23);
            this.progressBar_pack.TabIndex = 8;
            // 
            // button_pack_add
            // 
            this.button_pack_add.Location = new System.Drawing.Point(565, 273);
            this.button_pack_add.Name = "button_pack_add";
            this.button_pack_add.Size = new System.Drawing.Size(75, 23);
            this.button_pack_add.TabIndex = 7;
            this.button_pack_add.Text = "add";
            this.button_pack_add.UseVisualStyleBackColor = true;
            // 
            // button_pack_delete
            // 
            this.button_pack_delete.Location = new System.Drawing.Point(646, 273);
            this.button_pack_delete.Name = "button_pack_delete";
            this.button_pack_delete.Size = new System.Drawing.Size(75, 23);
            this.button_pack_delete.TabIndex = 6;
            this.button_pack_delete.Text = "delete";
            this.button_pack_delete.UseVisualStyleBackColor = true;
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Location = new System.Drawing.Point(30, 431);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(87, 17);
            this.label3.TabIndex = 5;
            this.label3.Text = "pack process:";
            // 
            // button_pack_select
            // 
            this.button_pack_select.Location = new System.Drawing.Point(646, 322);
            this.button_pack_select.Name = "button_pack_select";
            this.button_pack_select.Size = new System.Drawing.Size(75, 23);
            this.button_pack_select.TabIndex = 4;
            this.button_pack_select.Text = "select";
            this.button_pack_select.UseVisualStyleBackColor = true;
            // 
            // textBox_pack
            // 
            this.textBox_pack.Location = new System.Drawing.Point(33, 322);
            this.textBox_pack.Name = "textBox_pack";
            this.textBox_pack.Size = new System.Drawing.Size(607, 24);
            this.textBox_pack.TabIndex = 3;
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(30, 302);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(142, 17);
            this.label2.TabIndex = 2;
            this.label2.Text = "pack destination path:";
            // 
            // listView_pack
            // 
            this.listView_pack.Location = new System.Drawing.Point(33, 50);
            this.listView_pack.Name = "listView_pack";
            this.listView_pack.Size = new System.Drawing.Size(688, 217);
            this.listView_pack.TabIndex = 1;
            this.listView_pack.UseCompatibleStateImageBehavior = false;
            this.listView_pack.View = System.Windows.Forms.View.Details;
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(30, 30);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(111, 17);
            this.label1.TabIndex = 0;
            this.label1.Text = "pack source files:";
            // 
            // tabPage2
            // 
            this.tabPage2.Location = new System.Drawing.Point(4, 26);
            this.tabPage2.Name = "tabPage2";
            this.tabPage2.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage2.Size = new System.Drawing.Size(752, 507);
            this.tabPage2.TabIndex = 1;
            this.tabPage2.Text = "unpack";
            this.tabPage2.UseVisualStyleBackColor = true;
            // 
            // FormMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(96F, 96F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Dpi;
            this.ClientSize = new System.Drawing.Size(784, 561);
            this.Controls.Add(this.tab);
            this.DoubleBuffered = true;
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "FormMain";
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "package";
            this.FormClosed += new System.Windows.Forms.FormClosedEventHandler(this.FormMain_FormClosed);
            this.Load += new System.EventHandler(this.FormMain_Load);
            this.tab.ResumeLayout(false);
            this.tabPage1.ResumeLayout(false);
            this.tabPage1.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TabControl tab;
        private System.Windows.Forms.TabPage tabPage1;
        private System.Windows.Forms.TabPage tabPage2;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.ListView listView_pack;
        private System.Windows.Forms.TextBox textBox_pack;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.Button button_pack_select;
        private System.Windows.Forms.Button button_pack_add;
        private System.Windows.Forms.Button button_pack_delete;
        private System.Windows.Forms.Label label3;
        private System.Windows.Forms.Button button_pack_execute;
        private System.Windows.Forms.ProgressBar progressBar_pack;
        private System.Windows.Forms.ComboBox comboBox_pack;
        private System.Windows.Forms.Label label4;
    }
}

